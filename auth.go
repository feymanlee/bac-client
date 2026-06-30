package bac

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"
)

const defaultAuthVersion = "3"

func defaultNonce() int64 {
	return time.Now().UnixMilli()
}

// Sign returns the BAC Open API signature for the public query parameters.
func Sign(appKey, appSecret, authVersion, nonce string) string {
	plain := "appkey" + appKey + "auth_ver" + authVersion + "nonce" + nonce + appSecret
	sum := md5.Sum([]byte(plain))
	return hex.EncodeToString(sum[:])
}

func signedQuery(appKey, appSecret, authVersion, nonce string) url.Values {
	q := url.Values{}
	q.Set("appkey", appKey)
	q.Set("auth_ver", authVersion)
	q.Set("nonce", nonce)
	q.Set("s", Sign(appKey, appSecret, authVersion, nonce))
	return q
}

func signedQueryForNonce(appKey, appSecret, authVersion string, nonce int64) url.Values {
	return signedQuery(appKey, appSecret, authVersion, strconv.FormatInt(nonce, 10))
}
