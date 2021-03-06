package v1alpha1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var clusterUpdateRequestSchema *gojsonschema.Schema
var clusterDeleteRequestSchema *gojsonschema.Schema
var clusterDescribeRequestSchema *gojsonschema.Schema
var clusterListRequestSchema *gojsonschema.Schema
var clusterClientConfigRequestSchema *gojsonschema.Schema
var clusterReconfigureRequestSchema *gojsonschema.Schema
var clusterCreateRequestSchema *gojsonschema.Schema
var clusterApplyRequestSchema *gojsonschema.Schema
var clusterMetadataRequestSchema *gojsonschema.Schema

func init() {
	var err error
	clusterUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1alpha1ClusterSettings": {
      "properties": {
        "log_index_prefix": {
          "type": "string"
        },
        "log_storage_lifetime": {
          "title": "Number of secs logs will be stored in ElasticSearch",
          "type": "integer"
        },
        "monitoring_storage_lifetime": {
          "title": "Number of secs logs will be stored in InfluxDB",
          "type": "integer"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "default_access_level": {
      "title": "Default access level is to allow permission to the cluster\nwhen no Role matched for that specif user or group. This can\nset as\n  - kubernetes:team-admin\n  - kubernetes:cluster-admin\n  - kubernetes:admin\n  - kubernetes:editor\n  - kubernetes:viewer\n  - deny-access",
      "type": "string"
    },
    "do_not_delete": {
      "type": "boolean"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "settings": {
      "$ref": "#/definitions/v1alpha1ClusterSettings"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "delete_dynamic_volumes": {
      "type": "boolean"
    },
    "force": {
      "type": "boolean"
    },
    "keep_lodabalancers": {
      "type": "boolean"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "release_reserved_ip": {
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "status": {
      "items": {
        "type": "string"
      },
      "type": "array"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterClientConfigRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterReconfigureRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "apply_to_master": {
      "type": "boolean"
    },
    "count": {
      "type": "integer"
    },
    "kubernetes_version": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "sku": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1alpha1InstanceGroup": {
      "properties": {
        "count": {
          "type": "integer"
        },
        "sku": {
          "type": "string"
        },
        "use_spot_instances": {
          "type": "boolean"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "credential_uid": {
      "type": "string"
    },
    "default_access_level": {
      "title": "Default access level is to allow permission to the cluster\nwhen no Role matched for that specif user or group. This can\nset as\n  - kubernetes:team-admin\n  - kubernetes:cluster-admin\n  - kubernetes:admin\n  - kubernetes:editor\n  - kubernetes:viewer\n  - deny-access\nIf not set this will set \"\"",
      "type": "string"
    },
    "do_not_delete": {
      "type": "boolean"
    },
    "gce_project": {
      "type": "string"
    },
    "kubernetes_version": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "node_groups": {
      "items": {
        "$ref": "#/definitions/v1alpha1InstanceGroup"
      },
      "type": "array"
    },
    "provider": {
      "type": "string"
    },
    "zone": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterApplyRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterMetadataRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *ClusterUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterUpdateRequest) IsRequest() {}

func (m *ClusterDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDeleteRequest) IsRequest() {}

func (m *ClusterDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDescribeRequest) IsRequest() {}

func (m *ClusterListRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterListRequest) IsRequest() {}

func (m *ClusterClientConfigRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterClientConfigRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterClientConfigRequest) IsRequest() {}

func (m *ClusterReconfigureRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterReconfigureRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterReconfigureRequest) IsRequest() {}

func (m *ClusterCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterCreateRequest) IsRequest() {}

func (m *ClusterApplyRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterApplyRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterApplyRequest) IsRequest() {}

func (m *ClusterMetadataRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterMetadataRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterMetadataRequest) IsRequest() {}
