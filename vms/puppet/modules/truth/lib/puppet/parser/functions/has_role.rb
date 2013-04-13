module Puppet::Parser::Functions
  newfunction(:has_role, :type => :rvalue) do |args|
    if args.is_a? String
      args = [args]
    end
    role = args[0]
    roles = lookupvar("role")
    return false if roles == :undefined
    return roles.split(',').include? role
end

end

