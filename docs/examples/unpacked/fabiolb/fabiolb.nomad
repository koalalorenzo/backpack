
job {{ quote .job_name }} {
  {{ if .region }}region = {{ .region }}{{ end }}
  datacenters = {{ toJson .datacenters }}
  type = {{ quote .type }}

  group "fabio" {
    count = {{ .count | default 1 }}

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

    task "fabio" {

      {{- if eq .driver "docker" }}
      driver = "docker"
      
      config {
        image = "{{ .docker.image }}:{{ .docker.tag }}"
        {{- if .services.network_mode }}
        network_mode = {{ quote .services.network_mode }}
        {{- end }}
        port_map {
          proxy = 9999
          {{- if .services.ui.enable }}
          ui = 9998 
          {{- end }}
        }
      }
      
      {{ else }}
      {{/* In case of exec or raw_exec */}}
      driver = {{ quote .driver }}
      {{- with .exec }}
      artifact {
        source = "https://github.com/fabiolb/fabio/releases/download/v{{ .version }}/fabio-{{ .fullVersion }}-${attr.kernel.name}_${attr.cpu.arch}"
      }
      config {
        command = "./fabio-{{ .fullVersion }}-${attr.kernel.name}_${attr.cpu.arch}"
      }
      {{/* Close With .exec */}}
      {{- end }}

      {{/* Close With driver check */}}
      {{- end }}

      resources {
      {{- if .resources }}
        cpu    = {{ .resources.cpu }}
        memory = {{ .resources.memory }}
      {{- end }}
      {{- if eq .driver "docker" }}
        network {
          port "proxy" {
            {{- if .services.proxy.port }}
            static = {{ .services.proxy.port }}
            {{- end }}
          }
          port "ui" {
            {{- if .services.ui.port }} 
            static = {{ .services.ui.port }}
            {{- end }}
          }
        }
      {{- end }}
      }

      env {
        FABIO_PROXY_ADDR = ":${NOMAD_PORT_proxy}"
        {{- if .services.ui.enable }}
        FABIO_UI_ADDR = ":${NOMAD_PORT_ui}"
        {{- end }}
      }
    }

    {{- if ne .driver "docker" }}
    network {
      mode = {{ quote .services.network_mode }}
      port "proxy" {
        {{- if .services.proxy.port }}
        static = {{ .services.proxy.port }}
        {{- end }}
      }
      {{- if .services.ui.enable }} 
      port "ui" {
        {{- if .services.ui.port }} 
        static = {{ .services.ui.port }}
        {{- end }}
      }
      {{- end }}
    }
    {{- end }}
  }
}