package gitkit

import "strings"

type protocolVersion int

const (
	protocolUnknownVersion protocolVersion = -1
	protocol_v0            protocolVersion = 0
	protocol_v1            protocolVersion = 1
	protocol_v2            protocolVersion = 2
)

func parseProtocolVersion(value string) protocolVersion {
	switch value {
	case "0":
		return protocol_v0
	case "1":
		return protocol_v1
	case "2":
		return protocol_v2
	default:
		return protocolUnknownVersion
	}
}

func determineProtocolVersion(gitProtocol string) protocolVersion {
	version := protocol_v0

	if len(gitProtocol) != 0 {
		list := strings.Split(gitProtocol, ":")

		for _, item := range list {
			if strings.HasPrefix(item, "version=") {
				v := parseProtocolVersion(strings.TrimPrefix(item, "version="))
				if v > version {
					version = v
				}
			}
		}
	}

	return version
}
