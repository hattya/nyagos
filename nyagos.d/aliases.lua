nyagos.alias.ls='ls -oF $*'
nyagos.alias.lua_e=function(args) assert(load(args[1]))() end
nyagos.alias.lua_f=function(args)
    local path=table.remove(args,1)
    assert(loadfile(path))(args)
end
nyagos.alias["for"]='%COMSPEC% /c "@set PROMPT=$G & @for $*"'
nyagos.alias.kill = function(args)
    local command="taskkill.exe"
    for i=1,#args do
        if args[i] == "-f" then
            command="taskkill.exe /F"
        else
            nyagos.exec(command .. " /PID " .. args[i])
        end
    end
end
nyagos.alias.killall = function(args)
    local command="taskkill.exe"
    for i=1,#args do
        if args[i] == "-f" then
            command="taskkill.exe /F"
        else
            nyagos.exec(command .. " /IM " .. args[i])
        end
    end
end
