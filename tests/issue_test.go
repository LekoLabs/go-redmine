package redmine_test

import (
	"encoding/json"
	"testing"

	gomine "github.com/LekoLabs/go-redmine"
	"github.com/stretchr/testify/assert"
)

func Test_UnmarshalIssueJson(t *testing.T) {
	var issue gomine.Issue
	var exampleJson string = exampleIssueStruct()
	err := json.Unmarshal([]byte(exampleJson), &issue)

	// Assert that there's no error during unmarshaling
	assert.Nil(t, err, "Error unmarshaling example Json to Issue struct: ", err)

	// Assert the basic fields
	assert.Equal(t, 1, issue.Id)
	assert.Equal(t, "Sample Issue", issue.Subject)
	assert.Equal(t, "This is a sample issue", issue.Description)
	assert.Equal(t, "Initial creation of the issue", issue.Notes)

	// Assert the nested IdName fields
	assert.NotNil(t, issue.Project)
	assert.Equal(t, 101, issue.Project.Id)
	assert.Equal(t, "Sample Project", issue.Project.Name)

	assert.NotNil(t, issue.Tracker)
	assert.Equal(t, 102, issue.Tracker.Id)
	assert.Equal(t, "Bug", issue.Tracker.Name)

	assert.NotNil(t, issue.Status)
	assert.Equal(t, 103, issue.Status.Id)
	assert.Equal(t, "Open", issue.Status.Name)

	assert.NotNil(t, issue.Priority)
	assert.Equal(t, 104, issue.Priority.Id)
	assert.Equal(t, "High", issue.Priority.Name)

	assert.NotNil(t, issue.Author)
	assert.Equal(t, 105, issue.Author.Id)
	assert.Equal(t, "John Doe", issue.Author.Name)

	assert.NotNil(t, issue.FixedVersion)
	assert.Equal(t, 106, issue.FixedVersion.Id)
	assert.Equal(t, "v1.0", issue.FixedVersion.Name)

	assert.NotNil(t, issue.AssignedTo)
	assert.Equal(t, 107, issue.AssignedTo.Id)
	assert.Equal(t, "Jane Smith", issue.AssignedTo.Name)

	assert.NotNil(t, issue.Category)
	assert.Equal(t, 108, issue.Category.Id)
	assert.Equal(t, "UI", issue.Category.Name)

	// Assert the custom fields
	assert.Len(t, issue.CustomFields, 1)
	customField := issue.CustomFields[0]
	assert.Equal(t, 1, customField.Id)
	assert.Equal(t, "Severity", customField.Name)
	assert.Equal(t, "How severe the issue is", customField.Description)
	assert.Equal(t, false, customField.Multiple)
	assert.Equal(t, "Critical", customField.Value)

	// Assert the uploads
	assert.Len(t, issue.Uploads, 1)
	upload := issue.Uploads[0]
	assert.Equal(t, "abc123", upload.Token)
	assert.Equal(t, "screenshot.png", upload.Filename)
	assert.Equal(t, "image/png", upload.ContentType)

	// Assert the journals
	assert.Len(t, issue.Journals, 1)
	journal := issue.Journals[0]
	assert.Equal(t, 1, journal.Id)
	assert.NotNil(t, journal.User)
	assert.Equal(t, 201, journal.User.Id)
	assert.Equal(t, "User1", journal.User.Name)
	assert.Equal(t, "Started working on the issue", journal.Notes)
	assert.Equal(t, "2024-07-11T10:15:00Z", journal.CreatedOn)
	assert.Len(t, journal.Details, 1)
	journalDetail := journal.Details[0]
	assert.Equal(t, "attr", journalDetail.Property)
	assert.Equal(t, "status", journalDetail.Name)
	assert.Equal(t, "New", journalDetail.OldValue)
	assert.Equal(t, "In Progress", journalDetail.NewValue)

	// Assert that there is nothing on the Extra field
	assert.Len(t, issue.Extra, 0, "Expected Extra field to contain nothing; found %d elements in it", len(issue.Extra))
}

func Test_UnmarshalIssueJsonWithExtraFields(t *testing.T) {
	var issue gomine.Issue
	var exampleJson string = exampleIssueStructWithExtraFields()
	err := json.Unmarshal([]byte(exampleJson), &issue)

	// Assert that there's no error during unmarshaling
	assert.Nil(t, err, "Error unmarshaling example Json to Issue struct: ", err)

	// Assert that the length of the extra field is correct
	assert.Len(t, issue.Extra, 3, "Expected to find %d extra fields, only found %d", 3, len(issue.Extra))

	// Assert the Extra fields
	assert.Len(t, issue.Extra, 3)
	assert.Equal(t, "extra_value_1", issue.Extra["extra_field_1"])
	assert.Equal(t, float64(42), issue.Extra["extra_field_2"]) // JSON numbers are unmarshaled as float64 by default
	extraField3, ok := issue.Extra["extra_field_3"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(82), extraField3["id"])
	assert.Equal(t, "single_value", extraField3["single_key"])
}

func exampleIssueStruct() string {
	return `{
	"id": 1,
	"subject": "Sample Issue",
	"description": "This is a sample issue",
	"project": {
		"id": 101,
		"name": "Sample Project"
	},
	"tracker": {
		"id": 102,
		"name": "Bug"
	},
	"parent": {
		"id": 2
	},
	"status": {
		"id": 103,
		"name": "Open"
	},
	"priority": {
		"id": 104,
		"name": "High"
	},
	"author": {
		"id": 105,
		"name": "John Doe"
	},
	"fixed_version": {
		"id": 106,
		"name": "v1.0"
	},
	"assigned_to": {
		"id": 107,
		"name": "Jane Smith"
	},
	"category": {
		"id": 108,
		"name": "UI"
	},
	"notes": "Initial creation of the issue",
	"status_date": "2024-07-11",
	"created_on": "2024-07-11T10:00:00Z",
	"updated_on": "2024-07-11T11:00:00Z",
	"start_date": "2024-07-12",
	"due_date": "2024-07-20",
	"closed_on": "2024-07-21",
	"custom_fields": [
		{
			"id": 1,
			"name": "Severity",
			"description": "How severe the issue is",
			"multiple": false,
			"value": "Critical"
		}
	],
	"uploads": [
		{
			"token": "abc123",
			"filename": "screenshot.png",
			"content_type": "image/png"
		}
	],
	"done_ratio": 50,
	"estimated_hours": 5,
	"journals": [
		{
			"id": 1,
			"user": {
				"id": 201,
				"name": "User1"
			},
			"notes": "Started working on the issue",
			"created_on": "2024-07-11T10:15:00Z",
			"details": [
				{
					"property": "attr",
					"name": "status",
					"old_value": "New",
					"new_value": "In Progress"
				}
			]
		}
	]
}
`
}

func exampleIssueStructWithExtraFields() string {
	return `{
	"id": 1,
	"subject": "Sample Issue",
	"description": "This is a sample issue",
	"project": {
		"id": 101,
		"name": "Sample Project"
	},
	"tracker": {
		"id": 102,
		"name": "Bug"
	},
	"parent": {
		"id": 2
	},
	"status": {
		"id": 103,
		"name": "Open"
	},
	"priority": {
		"id": 104,
		"name": "High"
	},
	"author": {
		"id": 105,
		"name": "John Doe"
	},
	"fixed_version": {
		"id": 106,
		"name": "v1.0"
	},
	"assigned_to": {
		"id": 107,
		"name": "Jane Smith"
	},
	"category": {
		"id": 108,
		"name": "UI"
	},
	"notes": "Initial creation of the issue",
	"status_date": "2024-07-11",
	"created_on": "2024-07-11T10:00:00Z",
	"updated_on": "2024-07-11T11:00:00Z",
	"start_date": "2024-07-12",
	"due_date": "2024-07-20",
	"closed_on": "2024-07-21",
	"custom_fields": [
		{
			"id": 1,
			"name": "Severity",
			"description": "How severe the issue is",
			"multiple": false,
			"value": "Critical"
		}
	],
	"uploads": [
		{
			"token": "abc123",
			"filename": "screenshot.png",
			"content_type": "image/png"
		}
	],
	"done_ratio": 50,
	"estimated_hours": 5,
	"journals": [
		{
			"id": 1,
			"user": {
				"id": 201,
				"name": "User1"
			},
			"notes": "Started working on the issue",
			"created_on": "2024-07-11T10:15:00Z",
			"details": [
				{
					"property": "attr",
					"name": "status",
					"old_value": "New",
					"new_value": "In Progress"
				}
			]
		}
	],
	"extra_field_1": "extra_value_1",
	"extra_field_2": 42,
	"extra_field_3": {
		"id": 82,
		"single_key": "single_value"
	}
}
`
}
