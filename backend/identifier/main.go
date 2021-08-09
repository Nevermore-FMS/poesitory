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

func ParseStringIdentifier(str string) (*PluginVersionIdentifier, error) {
	const (
		stateName int = iota
		stateVersion
		stateChannel
	)
	var state = stateName

	var name string = ""
	var version string = ""
	var channel string = ""

	for _, c := range str {
		switch {
		case state < stateVersion && c == '@':
			state = stateVersion
		case state < stateChannel && c == '#':
			state = stateChannel
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
		if ok, _ := regexp.MatchString("^\\d+.\\d+.\\d+$", version); !ok {
			return nil, errors.New("invalid version")
		}
		vs := strings.Split(version, ".")
		major, err := strconv.Atoi(vs[0])
		if err != nil {
			return nil, err
		}
		minor, err := strconv.Atoi(vs[1])
		if err != nil {
			return nil, err
		}
		patch, err := strconv.Atoi(vs[2])
		if err != nil {
			return nil, err
		}
		identifier.Version = &PluginSemVer{
			Major: major,
			Minor: minor,
			Patch: patch,
		}
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

func ConstructIdentifierString(identifier PluginVersionIdentifier, latestVersion *PluginSemVer) string {
	var str = ""

	str += identifier.Name

	if latestVersion != nil {
		if *identifier.Version == *latestVersion {
			identifier.Version = nil
		}
	}

	if identifier.Version != nil {
		str += "@"
		str += fmt.Sprintf("%d.%d.%d", identifier.Version.Major, identifier.Version.Minor, identifier.Version.Patch)
	}

	if identifier.Channel != "STABLE" {
		str += "#"
		str += identifier.Channel
	}

	return str
}
