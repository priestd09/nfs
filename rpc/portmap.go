package rpc

import (
	"fmt"
	"encoding/binary"
)

// PORTMAP
// RFC 1057 Section A.1

const (
	PMAP_PORT = 111
	PMAP_PROG = 100000
	PMAP_VERS = 2

	PMAPPROC_GETPORT = 3
	PMAPPROC_DUMP = 4

	IPPROTO_TCP = 6
	IPPROTO_UDP = 17
)

type Mapping struct {
	Prog uint32
	Vers uint32
	Prot uint32
	Port uint32
}

type Portmapper struct {
	*Client
}

func (p *Portmapper) Getport(mapping Mapping) (int, error) {
	type getport struct {
		Header
		Mapping
	}
	msg := &getport {
		Header {
                        Rpcvers: 2,
                        Prog: PMAP_PROG,
                        Vers: PMAP_VERS,
                        Proc: PMAPPROC_GETPORT,                        
                        Cred: AUTH_NULL,
                        Verf: AUTH_NULL,
                },
		mapping,
	}
	buf, err := p.Call(msg)
	return int(binary.BigEndian.Uint32(buf)), err
}

func (p *Portmapper) Dump() (error) {
	type dump struct {
		Header
	}
	msg := &dump {
		Header {
			Rpcvers: 2,
			Prog: PMAP_PROG,
			Vers: PMAP_VERS,
			Proc: PMAPPROC_DUMP,
			Cred: AUTH_NULL,
			Verf: AUTH_NULL,
		},
	}
	_, err := p.Call(msg)
	return err
}	

func DialPortmapper(net, host string) (*Portmapper, error) {
	client, err := DialTCP(net, fmt.Sprintf("%s:%d", host, PMAP_PORT))
	if err != nil {
		return nil, err
	}
	return &Portmapper{client}, nil
}