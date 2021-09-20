function envoy_on_request(request_handle)
  request_handle:logInfo("SaaS Route preprocessing")
  local headers = request_handle:headers()
  local path = headers:get(":path")
  local debug = headers:get("X-Debug")
  local authority = headers:get(":authority")
  local forwardHeader = headers:get("X-SpendIntel-Forward")
  local saasRoutes = {
    "/saas/"
  }

  if authority == "testarossa-sfc.ivanticlouddev.com" and forwardHeader == "enabled" then
    for _,route in pairs(saasRoutes) do
      local pattern = "^(" .. route  .. ")"
      headers:remove("x-spendintel-forward")
      if string.match(path:lower(), pattern) then
          request_handle:headers():add("X-Debug-SAAS-Trace-Pattern-tested", pattern);
          local newPath = path:gsub("/saas/", "/api/cluster-saas/")
          
          if debug ~= nil and debug == "true" then
          request_handle:headers():add("X-Debug-SAAS-Trace-Path", path);
          request_handle:headers():add("X-Debug-SAAS-host-replaced", "true");
          request_handle:headers():add("X-Debug-SAAS-Trace-Path-ReplacedTo", newPath);
          end
          
          request_handle:headers():replace(":path", newPath)
          break
      end
    end
    
    if debug ~= nil and debug == "true" then
      request_handle:headers():add("X-Debug-SAAS-routes", "request.rule.executed");
      request_handle:headers():add("X-debug-SAAS-host", authority);
    end
  end
end

-- Enabling cors with new custom header on all platforms
function envoy_on_response(response_handle)
  local headers = response_handle:headers()
  local cors = headers:get("access-control-allow-headers")
  if cors ~= "" or cors ~= nil  then
    headers:remove("access-control-allow-headers")
    local newCors = "" .. cors .. ",X-SpendIntel-Forward"
    headers:add("access-control-allow-headers", newCors)
  end
end
