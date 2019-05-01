package libvirt

const (
	DriverName    = "libvirt"
	DriverVersion = "0.11.0"

	connectionString = "qemu:///system"
	dnsmasqLeases    = "/var/lib/libvirt/dnsmasq/%s.leases"
	dnsmasqStatus    = "/var/lib/libvirt/dnsmasq/%s.status"
	DefaultMemory    = 8096
	DefaultCPUs      = 4
	DefaultNetwork   = "crc"
	DefaultCacheMode = "default"
	DefaultIOMode    = "threads"
	DefaultSSHUser   = "core"
	DefaultSSHPort   = 22
	DomainTemplate   = `<domain type='kvm' xmlns:qemu='http://libvirt.org/schemas/domain/qemu/1.0'>
  <name>{{ .DomainName }}</name>
  <memory unit='MB'>{{ .Memory }}</memory>
  <vcpu placement='static'>{{ .CPU }}</vcpu>
  <features><acpi/><apic/><pae/></features>
  <cpu mode='host-passthrough'></cpu>
  <os>
    <type arch='x86_64'>hvm</type>
    <boot dev='hd'/>
    <bootmenu enable='no'/>
  </os>
  <features>
    <acpi/>
    <apic/>
    <pae/>
  </features>
  <clock offset='utc'/>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <devices>
    <disk type='file' device='disk'>
      <driver name='qemu' type='qcow2' cache='{{ .CacheMode }}' io='{{ .IOMode }}' />
      <source file='{{ .DiskPath }}'/>
      <target dev='vda' bus='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x05' function='0x0'/>
    </disk>
    <graphics type='vnc' autoport='yes' listen='127.0.0.1'>
      <listen type='address' address='127.0.0.1'/>
    </graphics>
    <controller type='usb' index='0' model='piix3-uhci'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x01' function='0x2'/>
    </controller>
    <controller type='pci' index='0' model='pci-root'/>
    <controller type='virtio-serial' index='0'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x04' function='0x0'/>
    </controller>
    <interface type='network'>
      <mac address='52:fd:fc:07:21:82'/>
      <source network='{{.Network}}'/>
      <model type='virtio'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x03' function='0x0'/>
    </interface>
    <serial type='pty'>
      <target type='isa-serial' port='0'>
        <model name='isa-serial'/>
      </target>
    </serial>
    <console type='pty'>
      <target type='serial' port='0'/>
    </console>
    <channel type='pty'>
      <target type='virtio' name='org.qemu.guest_agent.0'/>
      <address type='virtio-serial' controller='0' bus='0' port='1'/>
    </channel>
    <input type='mouse' bus='ps2'/>
    <input type='keyboard' bus='ps2'/>
    <graphics type='spice' autoport='yes'>
      <listen type='address'/>
    </graphics>
    <video>
      <model type='cirrus' vram='16384' heads='1' primary='yes'/>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x02' function='0x0'/>
    </video>
    <memballoon model='virtio'>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x06' function='0x0'/>
    </memballoon>
    <rng model='virtio'>
      <backend model='random'>/dev/random</backend>
      <address type='pci' domain='0x0000' bus='0x00' slot='0x07' function='0x0'/>
    </rng>
  </devices>
</domain>`
)

