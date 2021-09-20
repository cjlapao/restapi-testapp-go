  
  local expectedAuthority = "local-cluster.interna1l"
  expectedAuthority = expectedAuthority:lower()
  local auhtority = "local-cluster.internal"
  auhtority = auhtority:lower()  
  local path = "/api/importsoussrces('f)"
  local lcmRoutes = {
    "/api/factsonlyinstallationcomparison",
    "/api/importinstances",
    "/api/importsources"
  }
  if auhtority == expectedAuthority then
    for index,route in pairs(lcmRoutes) do
      print(route)
      local pattern = "^(" .. route .. ")"
      print(pattern)
      if string.match(path:lower(), pattern) then
        print("got it")
        break
      end
    end  
  end

--   function envoy_on_request(request_handle)
--     request_handle:logInfo("LCM Route preprocessing"); 
--     local headers = request_handle:headers()
--     local path = headers:get(":path")
--     local debug = headers:get("X-Debug")
--       request_handle:headers():add("X-Test", "cannot");
--     lcmRoutes = {}
--     lcmRoutes["test"] = "me"
--       request_handle:headers():add("X-LCM-Trace-Pattern", lcmRoutes["test"]);
--     for _,route in pairs(lcmRoutes) do
--       request_handle:headers():add("X-LCM-Trace-Pattern", route);
--       local pattern = "^(" .. route  .. ")"

--       if string.match(path:lower(), pattern) then
--         local newPath = path:gsub("/banana", "/status")
--         request_handle:headers():add("X-LCM-Trace-Path", path);
--         request_handle:headers():add("X-debug-lcm-host-replaced", "true");
--         request_handle:headers():add("X-LCM-Trace-Path-ReplacedTo", newSegment);
--         request_handle:headers():replace(":path", newSegment)
        
--       end
--     end
    
--     if debug ~= nil and debug == "true" then
--       request_handle:headers():add("X-debug-lcm", "request.rule.executed");
--       request_handle:headers():add("X-debug-lcm-host", path);
--     end
-- end