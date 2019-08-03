package jenkins

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/MacoTasu/gopher/config"
)

type Jenkins struct {
	Config config.ConfData
}

type BuildParameters struct {
	TaskName   string
	Parameters url.Values
}

func (j *Jenkins) Build(bp BuildParameters) error {
	jenkinsUrl := j.Config.Jenkins.Url

	requestURL := fmt.Sprintf(
		"%s/%s/buildWithParameters?%s",
		jenkinsUrl,
		bp.TaskName,
		bp.Parameters.Encode(),
	)

	fmt.Println(requestURL)
	resp, err := http.Get(requestURL)
	if err != nil {
		return fmt.Errorf("jenkins request error: %s", err)
	}
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("jenkins request response error: %s", resp.Status)
	}

	return nil
}
