package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

const socks5Ver = 0x05
const cmdBind = 0x01
const atypeIPV4 = 0x01
const atypeHOST = 0x03
const atypeIPV6 = 0x04

func main() {
	serve, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	for {
		client, err := serve.Accept()
		if err != nil {
			log.Printf("%v", err)
			continue
		}

		go process(client)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	err := auth(reader, conn)

	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")

	err = connect(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}

}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	ver, err := reader.ReadByte()

	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}

	if ver != socks5Ver {
		return fmt.Errorf("not support ver")
	}

	methodSize, err := reader.ReadByte()

	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}

	method := make([]byte, methodSize)

	_, err = io.ReadFull(reader, method)

	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}

	log.Println("ver", ver, "method", method)

	_, err = conn.Write([]byte{socks5Ver, 0x00})

	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}

	return nil

}

func connect(reader *bufio.Reader, conn net.Conn) (err error) {
	buf := make([]byte, 4)
	_, err = io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header fail")
	}
	ver, cmd, atyp := buf[0], buf[1], buf[3]

	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	if cmd != cmdBind {
		return fmt.Errorf("not supported cmd:%v", cmd)
	}
	addr := ""

	switch atyp {
	case atypeIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read atyp failed:%w", err)
		}
		addr = fmt.Sprintf("%d.%d.%d.%d", buf[0], buf[1], buf[2], buf[3])

	case atypeHOST:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read hostSize failed:%w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed:%w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("IPv6: no supported yet")
	default:
		return errors.New("invalid atyp")

	}
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed:%w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))
	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()
	log.Println("dial", addr, port)

	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()

	<-ctx.Done()
	return nil

}
