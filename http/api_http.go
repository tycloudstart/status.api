package http

import (
	"encoding/json"
	"github.com/tycloudstart.com/status.api/common/model"
	"github.com/tycloudstart.com/status.api/common/utils"
	"github.com/tycloudstart.com/status.api/proc"
	"net/http"
	"strings"
)

func configApiHttpRoutes() {
	// judge.status
	http.HandleFunc("/api/judge/status/", func(w http.ResponseWriter, r *http.Request) {
		// statistics
		proc.JudgeStatusQuery.Incr()

		urlParam := r.URL.Path[len("/api/judge/status/"):]
		args := strings.Split(urlParam, "/")

		argsLen := len(args)
		uuid := args[0]
		endpoint := args[1]
		metric := args[2]
		tags := make(map[string]string)
		if argsLen > 3 {
			tagVals := strings.Split(args[3], ",")
			for _, tag := range tagVals {
				tagPairs := strings.Split(tag, "=")
				if len(tagPairs) == 2 {
					tags[tagPairs[0]] = tagPairs[1]
				}
			}
		}
		RenderDataJson(w, model.NewJudgeStatus(uuid, endpoint, utils.Counter(metric, tags), 0))
	})

	// judge.statuses
	http.HandleFunc("/api/judge/statuses", func(w http.ResponseWriter, req *http.Request) {
		// statistics
		proc.JudgeStatusesQuery.Incr()

		if req.ContentLength == 0 {
			RenderErrorJson(w, "blank body")
			return
		}

		decoder := json.NewDecoder(req.Body)
		var queries []*model.JudgeStatusQuery
		err := decoder.Decode(&queries)
		if err != nil {
			RenderErrorJson(w, "decode error")
			return
		}

		statuses := make([]*model.JudgeStatus, 0)
		for _, query := range queries {
			statuses = append(statuses, model.NewJudgeStatus(query.Uuid, query.Endpoint, query.Counter, 0))
		}
		RenderDataJson(w, statuses)
	})
}
