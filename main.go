package main

import (
	_ "github.com/go-chi/chi"
	_ "github.com/golang/protobuf/proto"
	_ "github.com/kamilsk/go-kit/pkg/fn"
	_ "github.com/lib/pq"
	_ "github.com/mailru/easyjson"
	_ "github.com/pkg/errors"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/spf13/cobra"
	_ "github.com/spf13/viper"
	_ "github.com/stretchr/testify"
	_ "google.golang.org/grpc"
	_ "gopkg.in/yaml.v2"
)

func main() {}
