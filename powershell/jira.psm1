#
# jira powershell api
# 
# keep in mind that some of the process functions are jira setup specific and
# probably need to be readjusted for different projects and instances
#

# jira username
Set-Variable -Name "JIRA_USERNAME" -value "<TODO>" -scope global
# base64 encoded username:password
Set-Variable -Name "JIRA_BASIC_AUTH" -value "<TODO>" -scope global
# jira host
Set-Variable -Name "JIRA_HOST" -value "http://my.jira.com" -scope global

Set-Variable -Name "JIRA_REST_URL" -value "$JIRA_HOST/rest/api/2" -scope global
Set-Variable -Name "JIRA_AGILE_URL" -value "$JIRA_HOST/rest/agile/1.0" -scope global

# the id of your main jira board that you are using
Set-Variable -Name "JIRA_BOARD_ID" -value "10000" -scope global

function jiraIssueSetInProgress {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$status = $(jiraIssueGetStatus $JIRA_ISSUE_KEY)

		if ($status -EQ 'To Do') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Backlog
				jiraIssueSetTo $JIRA_ISSUE_KEY Progress
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "In Progress.")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Backlog') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Progress
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "In Progress.")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Progress') {
			echo "$(bold $JIRA_ISSUE_KEY) is already $(green "In Progress.")"
		}
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueSetInReview {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$status = $(jiraIssueGetStatus $JIRA_ISSUE_KEY)

		if ($status -EQ 'To Do') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Backlog
				jiraIssueSetTo $JIRA_ISSUE_KEY Progress
				jiraIssueSetTo $JIRA_ISSUE_KEY Review
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "In Review.")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Backlog') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Progress
				jiraIssueSetTo $JIRA_ISSUE_KEY Review
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "In Review.")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Progress') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Review
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "In Review.")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Review') {
			echo "$(bold $JIRA_ISSUE_KEY) is already $(green "In Review.")"
		}
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueSetInAwaitingTesting {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$status = $(jiraIssueGetStatus $JIRA_ISSUE_KEY)

		if ($status -EQ 'To Do') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Backlog
				jiraIssueSetTo $JIRA_ISSUE_KEY Progress
				jiraIssueSetTo $JIRA_ISSUE_KEY Review
				jiraIssueSetTo $JIRA_ISSUE_KEY Awaiting
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "Awaiting testing")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Backlog') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Progress
				jiraIssueSetTo $JIRA_ISSUE_KEY Review
				jiraIssueSetTo $JIRA_ISSUE_KEY Awaiting
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "Awaiting testing")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Progress') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Review
				jiraIssueSetTo $JIRA_ISSUE_KEY Awaiting
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "Awaiting testing")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'In Review') {
			try {
				jiraIssueSetTo $JIRA_ISSUE_KEY Awaiting
				echo "Set $(bold $JIRA_ISSUE_KEY) from $(red $status) to $(green "Awaiting testing")"
			} catch {
				echo "Something went wrong while trying to update $JIRA_ISSUE_KEY."
			}
		} elseif ($status -EQ 'Awaiting testing') {
			echo "$(bold $JIRA_ISSUE_KEY) is already $(green "Awaiting testing")"
		}
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueGetTransitions {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY/transitions"
		$RESULT = (Invoke-RestMethod -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json").transitions
		return $RESULT
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

# https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-rest-agile-1-0-board-boardid-sprint-get
function jiraBoardGetActiveSprint {
	if ($args.length -eq 1) {
		$BOARD_ID = $args[0]
		$URL = "$JIRA_AGILE_URL/board/$BOARD_ID/sprint?state=active"
		$RESULT = (Invoke-RestMethod -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json").values[0]
		return $RESULT
	} else {
		echo "requires one argument (board id). for example '1470'."
	}
}

# https://developer.atlassian.com/cloud/jira/software/rest/api-group-sprint/#api-rest-agile-1-0-sprint-sprintid-issue-post
function jiraSprintMoveTicket {
	if ($args.length -eq 2) {
		$SPRINT_ID = $args[0]
		$ISSUE_KEY = $args[1]
		$URL = "$JIRA_AGILE_URL/sprint/$SPRINT_ID/issue"
		$BODY = @{ 'issues' = @($ISSUE_KEY) } | ConvertTo-Json
		$RESULT = (Invoke-RestMethod -Method POST -Body $BODY -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json")
		return $RESULT
	} else {
		echo "requires two arguments (sprint id, issue key). for example '4846 D602451-658'."
	}
}

function jiraIssuesAssignedToMe {
	$URL = "$JIRA_REST_URL/search?jql=assignee=currentuser()&fields=status,summary"
	$RESULT = (Invoke-RestMethod -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json")
	return $RESULT.issues
}

function jiraIssueMoveToLatestSprint {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$SPRINT = $(jiraBoardGetActiveSprint $JIRA_BOARD_ID)
		$RESULT = jiraSprintMoveTicket $SPRINT.id $JIRA_ISSUE_KEY
		echo "Moved $(bold $JIRA_ISSUE_KEY) to sprint $(green $SPRINT.name)"
		return $RESULT
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueGetStatus {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY"
		$RESULT = (Invoke-RestMethod -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json").fields.status.name
		return $RESULT
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueDoTransition {
	if ($args.length -eq 2) {
		$JIRA_ISSUE_KEY = $args[0]
		$TRANSITION_ID = $args[1]
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY/transitions"
		$BODY = @{ 'transition' = @{ 'id' = $TRANSITION_ID } } | ConvertTo-Json
		$RESULT = (Invoke-RestMethod -Method POST -Body $BODY -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json")
	} else {
		echo "requires two arguments (jira issue key, transition id). for example 'JIRA-1173 51'."
	}
}

function jiraIssueUnassign {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY/assignee"
		$BODY = @{ 'name' = [wmi] $null } | ConvertTo-Json
		try {
			$RESULT = (Invoke-RestMethod -Method PUT -Body $BODY -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json")
			echo "$(green "Unassigned") $(bold $JIRA_ISSUE_KEY)."
		} catch {
			echo $(bold $(red "Problems while trying to unassign the ticket! Please check the issue yourself."))
		}
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueLogWork {
	if ($args.length -eq 3) {
		$JIRA_ISSUE_KEY = $args[0]
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY/worklog"
		$BODY = @{ 'comment' = $args[2]; 'timeSpent' = $args[1] } | ConvertTo-Json

		try {
			$RESULT = (Invoke-RestMethod -Method POST -Body $BODY -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json")
			echo $(bold $(green "Hours logged!"))
		} catch {
			echo $(bold $(red "Problems while logging work hours! Please check the issue yourself."))
		}
	} else {
		echo "requires three arguments (jira issue key, time spent and comment). for example 'JIRA-1173' '2h' 'implement rest api'."
	}
}

function jiraIssueAssignToMe {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY/assignee"
		$BODY = @{ 'name' = $JIRA_USERNAME } | ConvertTo-Json
		$RESULT = (Invoke-RestMethod -Method PUT -Body $BODY -Headers @{ 'Authorization' = "Basic $JIRA_BASIC_AUTH" } -Uri "$URL" -ContentType "application/json")
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueSetTo {
	if ($args.length -eq 2) {
		$JIRA_ISSUE_KEY = $args[0]
		$JIRA_ISSUE_STATUS = $args[1]
		$TRANSITION_ID = $($(jiraIssueGetTransitions $JIRA_ISSUE_KEY) | Where-Object -Property name -CMatch $JIRA_ISSUE_STATUS).id
		$URL = "$JIRA_REST_URL/issue/$JIRA_ISSUE_KEY/transitions"
		jiraIssueDoTransition $JIRA_ISSUE_KEY $TRANSITION_ID
	} else {
		echo "requires two argument (jira issue key, status [string]). for example 'JIRA-1173' 'Backlog'."
	}
}

function jiraIssueTestedOnDev {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		jiraIssueLogWork $JIRA_ISSUE_KEY 15m "test on dev"
		jiraIssueSetInAwaitingTesting $JIRA_ISSUE_KEY
		jiraIssueUnassign $JIRA_ISSUE_KEY
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}

function jiraIssueNotReproduceable {
	if ($args.length -eq 1) {
		$JIRA_ISSUE_KEY = $args[0]
		jiraIssueLogWork $JIRA_ISSUE_KEY 15m "trying to reproduce"
		jiraIssueAssignToMe $JIRA_ISSUE_KEY
		jiraIssueMoveToLatestSprint $JIRA_ISSUE_KEY
		jiraIssueSetInAwaitingTesting $JIRA_ISSUE_KEY
		jiraIssueUnassign $JIRA_ISSUE_KEY
	} else {
		echo "requires one argument (jira issue key). for example 'JIRA-1173'."
	}
}