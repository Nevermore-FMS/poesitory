package identifier

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PluginVersionIdentifier struct {
	Name    string
	Version *PluginSemVer
	Channel string
}

type PluginSemVer struct {
	Major int
	Minor int
	Patch int
}

// Format: plugin-name [ "#" CHANNEL ] [ "@" version ]
func ParseStringIdentifier(str string) (*PluginVersionIdentifier, error) {
	const (
		stateName int = iota
		stateChannel
		stateVersion
	)
	var state = stateName

	var name string = ""
	var channel string = ""
	var version string = ""

	for _, c := range str {
		switch {
		case state < stateChannel && c == '#':
			state = stateChannel
		case state < stateVersion && c == '@':
			state = stateVersion
		case state == stateName:
			name += string(c)
		case state == stateVersion:
			version += string(c)
		case state == stateChannel:
			channel += string(c)
		}
	}

	identifier := PluginVersionIdentifier{}

	if len(name) < 1 {
		return nil, errors.New("identifier must contain name")
	}
	identifier.Name = name

	if len(version) > 0 {
		semVer, err := ParseVersion(version)
		if err != nil {
			return nil, err
		}
		identifier.Version = &semVer
	}

	if len(channel) > 0 {
		if ok, _ := regexp.MatchString("^[A-Z]+$", channel); !ok {
			return nil, errors.New("invalid channel")
		}
		identifier.Channel = channel
	} else {
		identifier.Channel = "STABLE"
	}

	return &identifier, nil
}

func ParseVersion(vstring string) (PluginSemVer, error) {
	if ok, _ := regexp.MatchString("^\\d+.\\d+.\\d+$", vstring); !ok {
		return PluginSemVer{}, errors.New("invalid version")
	}
	vs := strings.Split(vstring, ".")
	major, err := strconv.Atoi(vs[0])
	if err != nil {
		return PluginSemVer{}, err
	}
	minor, err := strconv.Atoi(vs[1])
	if err != nil {
		return PluginSemVer{}, err
	}
	patch, err := strconv.Atoi(vs[2])
	if err != nil {
		return PluginSemVer{}, err
	}
	return PluginSemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func ConstructIdentifierString(identifier PluginVersionIdentifier, latestVersion *PluginSemVer) string {
	var str = ""

	str += identifier.Name

	if latestVersion != nil {
		if *identifier.Version == *latestVersion {
			identifier.Version = nil
		}
	}

	if identifier.Channel != "STABLE" {
		str += "#"
		str += identifier.Channel
	}

	if identifier.Version != nil {
		str += "@"
		str += fmt.Sprintf("%d.%d.%d", identifier.Version.Major, identifier.Version.Minor, identifier.Version.Patch)
	}

	return str
}
