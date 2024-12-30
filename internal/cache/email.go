package cache

import "context"

const emailPrifix = "email:"

func (r *RedisCache) GetEmailCode(ctx context.Context, email string) (string, error) {
	emailCode := r.client.Get(ctx, emailPrifix+email).Val()
	return emailCode, nil
}

func (r *RedisCache) SetEmailCode(ctx context.Context, email, emailCode string) error {
	return r.client.Set(ctx, emailPrifix+email, emailCode, r.emailCodeDuration).Err()
}
