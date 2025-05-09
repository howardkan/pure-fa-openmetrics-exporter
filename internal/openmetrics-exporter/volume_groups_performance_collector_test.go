package collectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"

	client "purestorage/fa-openmetrics-exporter/internal/rest-client"
)

func TestVolumesGroupPerformanceCollector(t *testing.T) {
	ref, _ := os.ReadFile("../../test/data/volume_groups_performance.json")
	vers, _ := os.ReadFile("../../test/data/versions.json")
	var volumegroupsperf client.VolumeGroupsPerformanceList
	json.Unmarshal(ref, &volumegroupsperf)
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vgperfu := regexp.MustCompile(`^/api/([0-9]+.[0-9]+)?/volume-groups/performance$`)
		if r.URL.Path == "/api/api_version" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(vers))
		} else if vgperfu.MatchString(r.URL.Path) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(ref))
		}
	}))
	endp := strings.Split(server.URL, "/")
	e := endp[len(endp)-1]
	defer server.Close()
	want := make(map[string]bool)
	for _, p := range volumegroupsperf.Items {
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"qos_rate_limit_usec_per_mirrored_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.QosRateLimitUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"qos_rate_limit_usec_per_read_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.QosRateLimitUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"qos_rate_limit_usec_per_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.QosRateLimitUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"queue_usec_per_mirrored_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.QueueUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"queue_usec_per_read_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.QueueUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"queue_usec_per_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.QueueUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"san_usec_per_mirrored_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.SanUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"san_usec_per_read_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.SanUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"san_usec_per_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.SanUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_mirrored_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ServiceUsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_read_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ServiceUsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ServiceUsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_mirrored_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.UsecPerMirroredWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_read_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.UsecPerReadOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"usec_per_write_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.UsecPerWriteOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"service_usec_per_read_op_cache_reduction\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ServiceUsecPerReadOpCacheReduction)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"mirrored_write_bytes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.MirroredWriteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"read_bytes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ReadBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"write_bytes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.WriteBytesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"mirrored_writes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.MirroredWritesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"reads_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.ReadsPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"writes_per_sec\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.WritesPerSec)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_mirrored_write\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerMirroredWrite)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_op\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerOp)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_read\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerRead)] = true
		want[fmt.Sprintf("label:{name:\"dimension\" value:\"bytes_per_write\"} label:{name:\"name\" value:\"%s\"} gauge:{value:%g}", p.Name, p.BytesPerWrite)] = true
	}
	c := client.NewRestClient(e, "fake-api-token", "latest", "test-user-agent-string", "test-X-Request-Id-string", false, false)

	vgpc := NewVolumeGroupsPerformanceCollector(c)
	metricsCheck(t, vgpc, want)
}
