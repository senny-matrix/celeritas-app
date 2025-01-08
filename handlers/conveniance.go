package handlers

import (
	"context"
	"fmt"
	"github.com/senny-matrix/celeritas"
	"net/http"
)

func (h *Handlers) render(w http.ResponseWriter, r *http.Request, tmpl string, variables, data any) error {
	return h.App.Render.Page(w, r, tmpl, variables, data)
}

func (h *Handlers) sessionPut(ctx context.Context, key string, val any) {
	h.App.Session.Put(ctx, key, val)
}

func (h *Handlers) sessionHas(ctx context.Context, key string) bool {
	return h.App.Session.Exists(ctx, key)
}

func (h *Handlers) sessionGet(ctx context.Context, key string) any {
	return h.App.Session.Get(ctx, key)
}

func (h *Handlers) sessionRemove(ctx context.Context, key string) {
	h.App.Session.Remove(ctx, key)
}

func (h *Handlers) sessionRenew(ctx context.Context) error {
	return h.App.Session.RenewToken(ctx)
}

func (h *Handlers) sessionDestroy(ctx context.Context) error {
	return h.App.Session.Destroy(ctx)
}

func (h *Handlers) randomString(n int) string {
	return h.App.RandomString(n)
}

func (h *Handlers) encrypt(text string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}

	ecrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return ecrypted, nil
}

func (h *Handlers) decrypt(crypto string) (string, error) {
	enc := celeritas.Encryption{Key: []byte(h.App.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return decrypted, nil
}

func (h *Handlers) TestCrypto(w http.ResponseWriter, r *http.Request) {
	plainText := "Hello World"
	fmt.Fprintf(w, "Plain Text: %s\n", plainText)
	encrypted, err := h.encrypt(plainText)
	if err != nil {
		h.App.ErrorLog.Println("Error encrypting:", err)
		h.App.Error500(w, r)
		return
	}
	fmt.Fprintf(w, "Encrypted Text: %s\n", encrypted)

	decrypted, err := h.decrypt(encrypted)
	if err != nil {
		h.App.ErrorLog.Println("Error decrypting:", err)
		h.App.Error500(w, r)
		return
	}
	fmt.Fprintf(w, "Decrypted Text: %s\n", decrypted)
}
