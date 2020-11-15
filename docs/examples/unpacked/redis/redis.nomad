job {{ quote .job_name }} {
  {{ if .region }}region = {{ .region }}{{ end }}
  datacenters = {{ toJson .datacenters }}
  type = {{ quote .type }}
  
  {{- if .update.enable }}
  update {
    min_healthy_time = {{ quote .update.min_healthy_time }}
    healthy_deadline = {{ quote .update.healthy_deadline }}
    {{- if not (eq .update.strategy "rolling") }}
    auto_revert      = true
    auto_promote     = true
    {{- else }}
    max_parallel = {{ .update.max_parallel }}
    {{- end }}
    {{- if eq .update.strategy "blue_green" }}
    max_parallel = {{ .update.count | default 1 }}
    canary = {{ .update.count | default 1 }}
    {{- end }}
    {{- if eq .update.strategy "canary" }}
    canary = 1
    {{- end }}
  }
  {{- end }}
  
  group "redis" {
    count = {{ .count | default 1 }}

    ephemeral_disk {
      size = 300
    }

    task "redis" {
      {{- if eq .driver "docker" }}
      driver = "docker"

      config {
        image = "{{ .docker.image }}:{{ .docker.tag }}"
        port_map {
          db = 6379
        }
      }
      {{- end }}

      {{- if and .persistency.enabled (eq .persistency.type "ephemeral") }}
      ephemeral_disk {
        migrate = true
        size    = {{ .persistency.size }}
        sticky  = true
      }
      {{- end }}

      resources {
      {{- if .resources }}
        cpu    = {{ .resources.cpu }}
        memory = {{ .resources.memory }}
      {{- end }}

        network {
          mbits = 10
          port "db" {}
        }
      }

      {{- if .service.enable }}
      service {
        name = {{ quote .service.name }}
        {{- if .service.tags }}
        tags = {{ toJson .service.tags }}
        {{- end }}
        port = "db"

        check {
          name     = "alive"
          type     = "tcp"
          interval = "10s"
          timeout  = "3s"
        }
      }
      {{- end }}
    }
  }
}
