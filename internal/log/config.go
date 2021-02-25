// Copyright © 2021 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"emperror.dev/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config holds details necessary for logging.
type Config struct {
	// Format specifies the output log format.
	// Accepted values are: json, logfmt
	Format string `json:"format,omitempty" mapstructure:"format"`

	// Level is the minimum log level that should appear on the output.
	Level string `json:"level,omitempty" mapstructure:"level"`

	// NoColor makes sure that no log output gets colorized.
	NoColor bool `json:"noColor,omitempty" mapstructure:"noColor"`
}

// Validate validates the configuration.
func (c Config) Validate() (Config, error) {
	if c.Format == "" {
		c.Format = "logfmt"
	}

	if c.Format != "json" && c.Format != "logfmt" {
		return c, errors.New("invalid log format: " + c.Format)
	}

	return c, nil
}

func ConfigureLoggingFlags(v *viper.Viper, p *pflag.FlagSet) {
	v.SetDefault("log.format", "json")
	p.String("log.level", "info", "Log level")
	v.SetDefault("log.level", "info")
	v.RegisterAlias("log.noColor", "no_color")
}
