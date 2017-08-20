#!/bin/bash
#
# Slack library for bash scripts.
#

declare SLACKTOKEN="${SLACKTOKEN:-$(cat /etc/secrets/slack/token.asc)}"
declare SLACKURL="https://hooks.slack.com/services"

slacksend() {
	local message
	message="${1}"

	if [[ "$#" -ne 1 ]]; then
		echo "Usage is $0 'Message for slack'" >&2
		return 1
	fi

 	echo "Sending a message '${message}' to slack"
	local response
	response=$(curl -s -H 'Content-type: application/json' \
	                --data "{\"text\":\"$message\"}" \
			${SLACKURL}/${SLACKTOKEN})

    	if [[ "${response}" != "ok" ]]; then
	      echo "Bad Slack response: '${response}'" >&2
	      return 1
	fi
	return 0
}
