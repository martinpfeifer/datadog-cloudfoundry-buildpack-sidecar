# shellcheck shell=bash
DD_TAGS=$(jq '
  [
    "application_id:\(.application_id)",
    "name:\(.name)",
    "space_name:\(.space_name)",
    (.uris[] | "uri:\(.)"),
    "cf_instance_ip:\(env.CF_INSTANCE_IP)",
    "cf_instance_index:\(env.CF_INSTANCE_INDEX)",
    "cf_instance_guid:\(env.CF_INSTANCE_GUID)"
  ] | join(" ")
' <<< "$VCAP_APPLICATION")
export DD_TAGS
