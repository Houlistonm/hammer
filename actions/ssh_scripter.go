package actions

import (
	"fmt"

	"github.com/pivotal/pcf/lockfile"
)

type SSHScripter struct{}

func NewSSHScripter() SSHScripter {
	return SSHScripter{}
}

func (b SSHScripter) Generate(data lockfile.Lockfile) []string {
	sshCommand := fmt.Sprintf(`ssh -o IdentitiesOnly=yes -o StrictHostKeyChecking=no -i "${ssh_key_path}" -t ubuntu@"%s"`, data.OpsManager.IP.String())

	sshCommandLines := []string{
		fmt.Sprintf(`ssh_key_path=$(mktemp)`),
		fmt.Sprintf(`echo "%s" >"$ssh_key_path"`, data.OpsManager.PrivateKey),
		fmt.Sprintf(`trap 'rm -f ${ssh_key_path}' EXIT`),
		fmt.Sprintf(`chmod 0600 "${ssh_key_path}"`),
		fmt.Sprintf(`creds="$(om -t %s -k -u %s -p %s curl -s -p %s)"`, data.OpsManager.URL.String(), data.OpsManager.Username, data.OpsManager.Password, boshCredsPath),
		fmt.Sprintf(`bosh="$(echo "$creds" | jq -r .credential | tr ' ' '\n' | grep '=')"`),
		fmt.Sprintf(`echo "$bosh"`),
		fmt.Sprintf(`shell="/usr/bin/env $(echo $bosh | tr '\n' ' ') bash -l"`),
		fmt.Sprintf(`%s "$shell"`, sshCommand),
	}

	return sshCommandLines
}
