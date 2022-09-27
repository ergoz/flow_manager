package sqlstore

import (
	"fmt"
	"strings"

	"github.com/lib/pq"
	"github.com/webitel/flow_manager/model"
	"github.com/webitel/flow_manager/store"
)

type SqlEmailStore struct {
	SqlStore
}

func NewSqlEmailStore(sqlStore SqlStore) store.EmailStore {
	st := &SqlEmailStore{sqlStore}
	return st
}

func (s SqlEmailStore) ProfileTaskFetch(node string) ([]*model.EmailProfileTask, *model.AppError) {
	var tasks []*model.EmailProfileTask

	_, err := s.GetReplica().Select(&tasks, ` update call_center.cc_email_profile
 set last_activity_at = now(),
     state = 'active'
 where enabled and
       last_activity_at < now() - (fetch_interval || ' sec')::interval
returning id, ( extract(EPOCH from updated_at) * 1000)::int8 updated_at`)

	if err != nil {
		return nil, model.NewAppError("SqlEmailStore.ProfileTaskFetch", "store.sql_email.task_profiles.error", nil,
			err.Error(), extractCodeFromErr(err))
	}

	return tasks, nil
}

func (s SqlEmailStore) Save(domainId int64, m *model.Email) *model.AppError {
	id, err := s.GetMaster().SelectInt(`insert into call_center.cc_email ("from", "to", profile_id, subject, cc, body, direction, message_id, sender, reply_to,
                      in_reply_to, parent_id, html)
values (:From, :To, :ProfileId, :Subject, :Cc, :Body::text, :Direction, :MessageId, :Sender, :ReplyTo, :InReplyTo, (select m.id
                                                                                                       from call_center.cc_email m
                                                                                                       where m.in_reply_to = :MessageId limit 1), 
		:Html::text)
	   returning id
`, map[string]interface{}{
		"From":      pq.Array(m.From),
		"To":        pq.Array(m.To),
		"ProfileId": m.ProfileId,
		"Subject":   m.Subject,
		"Cc":        pq.Array(m.CC),
		"Body":      m.Body,
		"Direction": m.Direction,
		"MessageId": m.MessageId,
		"Sender":    pq.Array(m.Sender),
		"ReplyTo":   pq.Array(m.ReplyTo),
		"InReplyTo": m.InReplyTo,
		"Html":      m.HtmlBody,
	})

	if err != nil {
		return model.NewAppError("SqlEmailStore.Save", "store.sql_email.save.error", nil,
			err.Error(), extractCodeFromErr(err))
	}
	m.Id = id
	return nil
}

func (s SqlEmailStore) GetProfile(id int) (*model.EmailProfile, *model.AppError) {
	var profile *model.EmailProfile

	err := s.GetReplica().SelectOne(&profile, `
select t.id, t.name, t.host, t.login, t.password, t.mailbox, t.imap_port, t.smtp_port, (extract(EPOCH from t.updated_at) * 1000)::int8 updated_at , t.flow_id, t.domain_id
from call_center.cc_email_profile t
where t.id = :Id`, map[string]interface{}{
		"Id": id,
	})

	if err != nil {
		return nil, model.NewAppError("SqlEmailStore.GetProfile", "store.sql_email.get_profile.error", nil,
			err.Error(), extractCodeFromErr(err))
	}

	return profile, nil
}

func (s SqlEmailStore) SetError(profileId int, appErr *model.AppError) *model.AppError {
	_, err := s.GetMaster().Exec(`update call_center.cc_email_profile
set enabled = false,
    fetch_err = :Err
where id = :Id`, map[string]interface{}{
		"Id":  profileId,
		"Err": appErr.Error(),
	})

	if err != nil {
		return model.NewAppError("SqlEmailStore.SetError", "store.sql_email_profile.set_error.error", nil,
			err.Error(), extractCodeFromErr(err))
	}

	return nil
}

func (s SqlEmailStore) GerProperties(domainId int64, id *int64, messageId *string, mapRes model.Variables) (model.Variables, *model.AppError) {
	f := make([]string, 0)

	for k, v := range mapRes {
		var val = ""
		switch v {
		case "from", "to", "subject",
			"cc", "sender", "reply_to", "in_reply_to", "body", "html":
			val = fmt.Sprintf("\"%s\" as %s", v, pq.QuoteIdentifier(k))
		}

		f = append(f, val)
	}

	var t *properties

	err := s.GetReplica().SelectOne(&t, `select row_to_json(t) variables
from (
    select `+strings.Join(f, ", ")+`
    from (
        select
            id,
            array_to_string("from", ',') as from,
            array_to_string("to", ',') as to,
            subject,
            array_to_string("cc", ',') as cc,
            array_to_string("sender", ',') as sender,
            array_to_string("reply_to", ',') as reply_to,
            in_reply_to,
            body,
            html
        from call_center.cc_email e
        where (id = :Id or message_id = :MessageId)
            and exists(select 1 from call_center.cc_email_profile p where p.domain_id = :DomainId and p.id = e.profile_id)
        order by e.created_at desc
        limit 1
     ) t
 ) t`, map[string]interface{}{
		"DomainId":  domainId,
		"Id":        id,
		"MessageId": messageId,
	})

	if err != nil {
		return nil, model.NewAppError("SqlEmailStore.Get", "store.sql_email.get.app_error", nil, err.Error(), extractCodeFromErr(err))
	}

	return t.Variables, nil
}
