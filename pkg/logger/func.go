package logger

import (
	"context"
	"github.com/sirupsen/logrus"
)

func With(ctx context.Context, value logrus.Fields) context.Context {
	l := Get(ctx)
	l = l.WithFields(value)
	return Set(ctx, l)
}

func WithSessionID(ctx context.Context, sessionID string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"SessionID": sessionID,
	})
	return Set(ctx, l)
}

func WithRequestID(ctx context.Context, requestID string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"RequestID": requestID,
	})
	return Set(ctx, l)
}

func WithUserID(ctx context.Context, userID string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"UserID": userID,
	})
	return Set(ctx, l)
}

func WithURL(ctx context.Context, url string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"URL": url,
	})
	return Set(ctx, l)
}

func WithPassword(ctx context.Context, password string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"Password": password,
	})
	return Set(ctx, l)
}

func WithConfirmationCode(ctx context.Context, confirmationCode string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"ConfirmationCode": confirmationCode,
	})
	return Set(ctx, l)
}

func WithRestAuthData(ctx context.Context, email, phone string) context.Context {
	if email != "" {
		ctx = WithEmail(ctx, email)
	}

	if phone != "" {
		ctx = WithPhone(ctx, phone)
	}
	return ctx
}

func WithEmail(ctx context.Context, email string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"Email": email,
	})
	return Set(ctx, l)
}

func WithPhone(ctx context.Context, phone string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"Phone": phone,
	})
	return Set(ctx, l)
}

func WithGroupName(ctx context.Context, groupName string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"GroupName": groupName,
	})
	return Set(ctx, l)
}

func WithGroupID(ctx context.Context, groupID string) context.Context {
	l := Get(ctx)
	l = l.WithFields(logrus.Fields{
		"GroupID": groupID,
	})
	return Set(ctx, l)
}
