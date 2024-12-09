package gobom

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)

func (g *Generator) detectLicense(depName string) string {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}", depName)
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return "Unknown"
	}

	depPath := strings.TrimSpace(string(output))
	if depPath == "" {
		return "Unknown"
	}

	licensePatterns := []string{
		"LICENSE",
		"LICENSE.txt",
		"LICENSE.md",
		"COPYING",
		"COPYING.txt",
		"COPYING.md",
	}

	for _, pattern := range licensePatterns {
		matches, err := filepath.Glob(filepath.Join(depPath, pattern+"*"))
		if err != nil {
			continue
		}

		for _, match := range matches {
			content, err := ioutil.ReadFile(match)
			if err != nil {
				continue
			}

			license := identifyLicense(string(content))
			if license != "Unknown" {
				return license
			}
		}
	}

	return "Unknown"
}

func identifyLicense(content string) string {
	licenseIdentifiers := map[string][]string{
		"MIT": {
			"MIT License",
			"Permission is hereby granted, free of charge",
			"without restriction, including without limitation the rights",
		},
		"Apache-2.0": {
			"Apache License",
			"Version 2.0",
			"Licensed under the Apache License",
		},
		"GPL-3.0": {
			"GNU GENERAL PUBLIC LICENSE",
			"Version 3",
			"GNU General Public License version 3",
		},
		"GPL-2.0": {
			"GNU GENERAL PUBLIC LICENSE",
			"Version 2",
			"GNU General Public License version 2",
		},
		"BSD-3-Clause": {
			"Redistribution and use in source and binary forms",
			"3. Neither the name",
			"BSD 3-Clause License",
		},
		"BSD-2-Clause": {
			"Redistribution and use in source and binary forms",
			"2. Redistributions in binary form",
			"BSD 2-Clause License",
		},
		"LGPL-3.0": {
			"GNU LESSER GENERAL PUBLIC LICENSE",
			"Version 3",
		},
		"MPL-2.0": {
			"Mozilla Public License Version 2.0",
			"Mozilla Public License, version 2.0",
		},
		"ISC": {
			"ISC License",
			"Permission to use, copy, modify, and/or distribute this software",
		},
	}

	content = strings.ToUpper(content)
	for license, patterns := range licenseIdentifiers {
		matches := 0
		for _, pattern := range patterns {
			if strings.Contains(content, strings.ToUpper(pattern)) {
				matches++
			}
		}
		if matches == len(patterns) {
			return license
		}
	}

	spdxIdentifiers := []string{
		"SPDX-License-Identifier:",
		"SPDX-FileCopyrightText:",
	}
	for _, identifier := range spdxIdentifiers {
		if idx := strings.Index(content, strings.ToUpper(identifier)); idx != -1 {
			start := idx + len(identifier)
			end := strings.Index(content[start:], "\n")
			if end == -1 {
				end = len(content)
			} else {
				end = start + end
			}
			return strings.TrimSpace(content[start:end])
		}
	}

	return "Unknown"
}
