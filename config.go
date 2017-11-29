/* Copyright © Playground Global, LLC. All rights reserved. */

package mail

type SMTPConfig struct {
	Server   string
	Port     int
	User     string
	Password string
}

type TemplateConfig struct {
	Name        string
	File        string
	SenderEmail string
}

type ConfigType struct {
	SMTP         *SMTPConfig
	TemplateRoot string
	Templates    []*TemplateConfig
}

var Config ConfigType = ConfigType{
	&SMTPConfig{
		"smtp.gmail.com",
		25,
		"noreply@domain.tld",
		"Sekr1tPassw0rd",
	},
	"./mails",
	[]*TemplateConfig{
		{
			"package",
			"package.tmpl",
			"noreply+reception@domain.tld",
		},
	},
}
