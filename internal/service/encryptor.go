package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/kordape/redbrain-challenge/pkg/encryptorapi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var blacklisted = []string{
	"fast and furious",
}

type EncryptionService struct {
	encryptorapi.UnimplementedEncryptorServer
}

func (s *EncryptionService) Encrypt(ctx context.Context, req *encryptorapi.EncryptRequest) (*encryptorapi.EncryptResponse, error) {
	// check if input is in the blacklist
	// simple check if implemented - checking if the input contains a substring on the blacklist string
	for _, b := range blacklisted {
		if strings.Contains(strings.ToLower(req.Data), b) {
			return nil, status.Error(codes.InvalidArgument, "Requested to encrypt forbidden input")
		}
	}

	hasher := sha256.New()
	_, err := hasher.Write([]byte(req.Data))
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to encrypt")
	}

	return &encryptorapi.EncryptResponse{
		Encrypted: hex.EncodeToString(hasher.Sum(nil)),
	}, nil
}
