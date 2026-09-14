package models

import (
	"encoding/json"
	"testing"
)

func TestCommentNodeScheme_AppendNode(t *testing.T) {
	type fields struct {
		Version int
		Type    string
		Content []*CommentNodeScheme
		Text    string
		Attrs   map[string]interface{}
		Marks   []*MarkScheme
	}
	type args struct {
		node *CommentNodeScheme
	}
	testCases := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "when the parameters are correct",
			fields: fields{},
			args: args{
				node: &CommentNodeScheme{
					Type: "text",
				},
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			n := &CommentNodeScheme{
				Version: testCase.fields.Version,
				Type:    testCase.fields.Type,
				Content: testCase.fields.Content,
				Text:    testCase.fields.Text,
				Attrs:   testCase.fields.Attrs,
				Marks:   testCase.fields.Marks,
			}
			n.AppendNode(testCase.args.node)
		})
	}
}

func TestIssueCommentScheme_ParentID_Unmarshal(t *testing.T) {
	t.Run("reply comment has parentId", func(t *testing.T) {
		raw := `{"id":"17835150","parentId":"14650948","body":{"type":"doc"}}`
		var comment IssueCommentScheme
		if err := json.Unmarshal([]byte(raw), &comment); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if comment.ID != "17835150" {
			t.Errorf("ID = %q, want 17835150", comment.ID)
		}
		if comment.ParentID != "14650948" {
			t.Errorf("ParentID = %q, want 14650948", comment.ParentID)
		}
	})

	t.Run("top-level comment has no parentId", func(t *testing.T) {
		raw := `{"id":"14650948","body":{"type":"doc"}}`
		var comment IssueCommentScheme
		if err := json.Unmarshal([]byte(raw), &comment); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if comment.ParentID != "" {
			t.Errorf("ParentID = %q, want empty for top-level comment", comment.ParentID)
		}
	})

	t.Run("parentId round-trips through marshal", func(t *testing.T) {
		original := IssueCommentScheme{ID: "100", ParentID: "99"}
		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("marshal failed: %v", err)
		}
		var decoded IssueCommentScheme
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("unmarshal failed: %v", err)
		}
		if decoded.ParentID != "99" {
			t.Errorf("round-trip ParentID = %q, want 99", decoded.ParentID)
		}
	})
}
