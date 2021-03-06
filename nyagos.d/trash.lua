if not share.ole then
    local status
    status,share.ole = pcall(require,"nyole")
    if not status then
        share.ole = nil
    end
end
if share.ole then
    nyagos.alias.trash = function(args)
        local fsObj = share.ole.create_object_utf8("Scripting.FileSystemObject")
        local shellApp = share.ole.create_object_utf8("Shell.Application")
        local trashBox = shellApp:NameSpace(math.tointeger(10))
        if trashBox.MoveHere then
            if #args <= 0 then
                nyagos.writerr("Move files or directories to Windows Trashbox\n")
                nyagos.writerr("Usage: trash file(s)...\n")
                return
            end
            args = nyagos.glob(table.unpack(args))
            for i=1,#args do
                if fsObj:FileExists(args[i]) or fsObj:FolderExists(args[i]) then
                    trashBox:MoveHere(fsObj:GetAbsolutePathName(args[i]))
                else
                    nyagos.writerr(args[i]..": such a file or directory not found.\n")
                end
            end
        else
            nyagos.writerr("Warning: trash.lua requires nyaole.dll 0.0.0.5 or later\n")
        end
    end
end
