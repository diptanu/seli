require 'facter'

if File.exists?('/etc/custom_facts.facts')
  File.readlines('/etc/custom_facts.facts').each { |line|
    if line =~ /^(.+)=(.+)$/
      var = "easycheckout_" + $1.strip
      val = $2.strip

      Facter.add(var) do
        setcode { val }
      end
      
    end
	}
end

  
  
