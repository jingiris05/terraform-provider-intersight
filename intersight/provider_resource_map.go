package intersight

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GetResourceMapping() map[string]*schema.Resource {
	return map[string]*schema.Resource{
		"intersight_aaa_retention_policy":                                    resourceAaaRetentionPolicy(),
		"intersight_access_policy":                                           resourceAccessPolicy(),
		"intersight_adapter_config_policy":                                   resourceAdapterConfigPolicy(),
		"intersight_appliance_auto_rma_policy":                               resourceApplianceAutoRmaPolicy(),
		"intersight_appliance_backup":                                        resourceApplianceBackup(),
		"intersight_appliance_backup_policy":                                 resourceApplianceBackupPolicy(),
		"intersight_appliance_data_export_policy":                            resourceApplianceDataExportPolicy(),
		"intersight_appliance_device_claim":                                  resourceApplianceDeviceClaim(),
		"intersight_appliance_diag_setting":                                  resourceApplianceDiagSetting(),
		"intersight_appliance_remote_file_import":                            resourceApplianceRemoteFileImport(),
		"intersight_appliance_restore":                                       resourceApplianceRestore(),
		"intersight_asset_target":                                            resourceAssetTarget(),
		"intersight_bios_policy":                                             resourceBiosPolicy(),
		"intersight_boot_precision_policy":                                   resourceBootPrecisionPolicy(),
		"intersight_bulk_export":                                             resourceBulkExport(),
		"intersight_bulk_request":                                            resourceBulkRequest(),
		"intersight_capability_adapter_unit_descriptor":                      resourceCapabilityAdapterUnitDescriptor(),
		"intersight_capability_chassis_descriptor":                           resourceCapabilityChassisDescriptor(),
		"intersight_capability_chassis_manufacturing_def":                    resourceCapabilityChassisManufacturingDef(),
		"intersight_capability_cimc_firmware_descriptor":                     resourceCapabilityCimcFirmwareDescriptor(),
		"intersight_capability_equipment_physical_def":                       resourceCapabilityEquipmentPhysicalDef(),
		"intersight_capability_equipment_slot_array":                         resourceCapabilityEquipmentSlotArray(),
		"intersight_capability_fan_module_descriptor":                        resourceCapabilityFanModuleDescriptor(),
		"intersight_capability_fan_module_manufacturing_def":                 resourceCapabilityFanModuleManufacturingDef(),
		"intersight_capability_io_card_capability_def":                       resourceCapabilityIoCardCapabilityDef(),
		"intersight_capability_io_card_descriptor":                           resourceCapabilityIoCardDescriptor(),
		"intersight_capability_io_card_manufacturing_def":                    resourceCapabilityIoCardManufacturingDef(),
		"intersight_capability_port_group_aggregation_def":                   resourceCapabilityPortGroupAggregationDef(),
		"intersight_capability_psu_descriptor":                               resourceCapabilityPsuDescriptor(),
		"intersight_capability_psu_manufacturing_def":                        resourceCapabilityPsuManufacturingDef(),
		"intersight_capability_server_models_capability_def":                 resourceCapabilityServerModelsCapabilityDef(),
		"intersight_capability_server_schema_descriptor":                     resourceCapabilityServerSchemaDescriptor(),
		"intersight_capability_sioc_module_capability_def":                   resourceCapabilitySiocModuleCapabilityDef(),
		"intersight_capability_sioc_module_descriptor":                       resourceCapabilitySiocModuleDescriptor(),
		"intersight_capability_sioc_module_manufacturing_def":                resourceCapabilitySiocModuleManufacturingDef(),
		"intersight_capability_switch_capability":                            resourceCapabilitySwitchCapability(),
		"intersight_capability_switch_descriptor":                            resourceCapabilitySwitchDescriptor(),
		"intersight_capability_switch_manufacturing_def":                     resourceCapabilitySwitchManufacturingDef(),
		"intersight_certificatemanagement_policy":                            resourceCertificatemanagementPolicy(),
		"intersight_chassis_config_import":                                   resourceChassisConfigImport(),
		"intersight_chassis_profile":                                         resourceChassisProfile(),
		"intersight_comm_http_proxy_policy":                                  resourceCommHttpProxyPolicy(),
		"intersight_connectorpack_connector_pack_upgrade":                    resourceConnectorpackConnectorPackUpgrade(),
		"intersight_convergedinfra_health_check_definition":                  resourceConvergedinfraHealthCheckDefinition(),
		"intersight_crd_custom_resource":                                     resourceCrdCustomResource(),
		"intersight_deviceconnector_policy":                                  resourceDeviceconnectorPolicy(),
		"intersight_externalsite_authorization":                              resourceExternalsiteAuthorization(),
		"intersight_fabric_appliance_pc_role":                                resourceFabricAppliancePcRole(),
		"intersight_fabric_appliance_role":                                   resourceFabricApplianceRole(),
		"intersight_fabric_eth_network_control_policy":                       resourceFabricEthNetworkControlPolicy(),
		"intersight_fabric_eth_network_group_policy":                         resourceFabricEthNetworkGroupPolicy(),
		"intersight_fabric_eth_network_policy":                               resourceFabricEthNetworkPolicy(),
		"intersight_fabric_fc_network_policy":                                resourceFabricFcNetworkPolicy(),
		"intersight_fabric_fc_storage_role":                                  resourceFabricFcStorageRole(),
		"intersight_fabric_fc_uplink_pc_role":                                resourceFabricFcUplinkPcRole(),
		"intersight_fabric_fc_uplink_role":                                   resourceFabricFcUplinkRole(),
		"intersight_fabric_fcoe_uplink_pc_role":                              resourceFabricFcoeUplinkPcRole(),
		"intersight_fabric_fcoe_uplink_role":                                 resourceFabricFcoeUplinkRole(),
		"intersight_fabric_flow_control_policy":                              resourceFabricFlowControlPolicy(),
		"intersight_fabric_link_aggregation_policy":                          resourceFabricLinkAggregationPolicy(),
		"intersight_fabric_link_control_policy":                              resourceFabricLinkControlPolicy(),
		"intersight_fabric_multicast_policy":                                 resourceFabricMulticastPolicy(),
		"intersight_fabric_pc_operation":                                     resourceFabricPcOperation(),
		"intersight_fabric_port_mode":                                        resourceFabricPortMode(),
		"intersight_fabric_port_operation":                                   resourceFabricPortOperation(),
		"intersight_fabric_port_policy":                                      resourceFabricPortPolicy(),
		"intersight_fabric_server_role":                                      resourceFabricServerRole(),
		"intersight_fabric_switch_cluster_profile":                           resourceFabricSwitchClusterProfile(),
		"intersight_fabric_switch_control_policy":                            resourceFabricSwitchControlPolicy(),
		"intersight_fabric_switch_profile":                                   resourceFabricSwitchProfile(),
		"intersight_fabric_system_qos_policy":                                resourceFabricSystemQosPolicy(),
		"intersight_fabric_uplink_pc_role":                                   resourceFabricUplinkPcRole(),
		"intersight_fabric_uplink_role":                                      resourceFabricUplinkRole(),
		"intersight_fabric_vlan":                                             resourceFabricVlan(),
		"intersight_fabric_vsan":                                             resourceFabricVsan(),
		"intersight_fcpool_pool":                                             resourceFcpoolPool(),
		"intersight_firmware_bios_descriptor":                                resourceFirmwareBiosDescriptor(),
		"intersight_firmware_board_controller_descriptor":                    resourceFirmwareBoardControllerDescriptor(),
		"intersight_firmware_chassis_upgrade":                                resourceFirmwareChassisUpgrade(),
		"intersight_firmware_cimc_descriptor":                                resourceFirmwareCimcDescriptor(),
		"intersight_firmware_dimm_descriptor":                                resourceFirmwareDimmDescriptor(),
		"intersight_firmware_distributable":                                  resourceFirmwareDistributable(),
		"intersight_firmware_drive_descriptor":                               resourceFirmwareDriveDescriptor(),
		"intersight_firmware_driver_distributable":                           resourceFirmwareDriverDistributable(),
		"intersight_firmware_eula":                                           resourceFirmwareEula(),
		"intersight_firmware_gpu_descriptor":                                 resourceFirmwareGpuDescriptor(),
		"intersight_firmware_hba_descriptor":                                 resourceFirmwareHbaDescriptor(),
		"intersight_firmware_iom_descriptor":                                 resourceFirmwareIomDescriptor(),
		"intersight_firmware_mswitch_descriptor":                             resourceFirmwareMswitchDescriptor(),
		"intersight_firmware_nxos_descriptor":                                resourceFirmwareNxosDescriptor(),
		"intersight_firmware_pcie_descriptor":                                resourceFirmwarePcieDescriptor(),
		"intersight_firmware_psu_descriptor":                                 resourceFirmwarePsuDescriptor(),
		"intersight_firmware_sas_expander_descriptor":                        resourceFirmwareSasExpanderDescriptor(),
		"intersight_firmware_server_configuration_utility_distributable":     resourceFirmwareServerConfigurationUtilityDistributable(),
		"intersight_firmware_storage_controller_descriptor":                  resourceFirmwareStorageControllerDescriptor(),
		"intersight_firmware_switch_upgrade":                                 resourceFirmwareSwitchUpgrade(),
		"intersight_firmware_unsupported_version_upgrade":                    resourceFirmwareUnsupportedVersionUpgrade(),
		"intersight_firmware_upgrade":                                        resourceFirmwareUpgrade(),
		"intersight_hcl_hyperflex_software_compatibility_info":               resourceHclHyperflexSoftwareCompatibilityInfo(),
		"intersight_hyperflex_app_catalog":                                   resourceHyperflexAppCatalog(),
		"intersight_hyperflex_auto_support_policy":                           resourceHyperflexAutoSupportPolicy(),
		"intersight_hyperflex_capability_info":                               resourceHyperflexCapabilityInfo(),
		"intersight_hyperflex_cluster_backup_policy":                         resourceHyperflexClusterBackupPolicy(),
		"intersight_hyperflex_cluster_backup_policy_deployment":              resourceHyperflexClusterBackupPolicyDeployment(),
		"intersight_hyperflex_cluster_network_policy":                        resourceHyperflexClusterNetworkPolicy(),
		"intersight_hyperflex_cluster_profile":                               resourceHyperflexClusterProfile(),
		"intersight_hyperflex_cluster_replication_network_policy":            resourceHyperflexClusterReplicationNetworkPolicy(),
		"intersight_hyperflex_cluster_replication_network_policy_deployment": resourceHyperflexClusterReplicationNetworkPolicyDeployment(),
		"intersight_hyperflex_cluster_storage_policy":                        resourceHyperflexClusterStoragePolicy(),
		"intersight_hyperflex_ext_fc_storage_policy":                         resourceHyperflexExtFcStoragePolicy(),
		"intersight_hyperflex_ext_iscsi_storage_policy":                      resourceHyperflexExtIscsiStoragePolicy(),
		"intersight_hyperflex_feature_limit_external":                        resourceHyperflexFeatureLimitExternal(),
		"intersight_hyperflex_feature_limit_internal":                        resourceHyperflexFeatureLimitInternal(),
		"intersight_hyperflex_health_check_definition":                       resourceHyperflexHealthCheckDefinition(),
		"intersight_hyperflex_health_check_package_checksum":                 resourceHyperflexHealthCheckPackageChecksum(),
		"intersight_hyperflex_hxdp_version":                                  resourceHyperflexHxdpVersion(),
		"intersight_hyperflex_local_credential_policy":                       resourceHyperflexLocalCredentialPolicy(),
		"intersight_hyperflex_node_config_policy":                            resourceHyperflexNodeConfigPolicy(),
		"intersight_hyperflex_node_profile":                                  resourceHyperflexNodeProfile(),
		"intersight_hyperflex_proxy_setting_policy":                          resourceHyperflexProxySettingPolicy(),
		"intersight_hyperflex_server_firmware_version":                       resourceHyperflexServerFirmwareVersion(),
		"intersight_hyperflex_server_firmware_version_entry":                 resourceHyperflexServerFirmwareVersionEntry(),
		"intersight_hyperflex_server_model":                                  resourceHyperflexServerModel(),
		"intersight_hyperflex_service_auth_token":                            resourceHyperflexServiceAuthToken(),
		"intersight_hyperflex_software_distribution_component":               resourceHyperflexSoftwareDistributionComponent(),
		"intersight_hyperflex_software_distribution_entry":                   resourceHyperflexSoftwareDistributionEntry(),
		"intersight_hyperflex_software_distribution_version":                 resourceHyperflexSoftwareDistributionVersion(),
		"intersight_hyperflex_software_version_policy":                       resourceHyperflexSoftwareVersionPolicy(),
		"intersight_hyperflex_sys_config_policy":                             resourceHyperflexSysConfigPolicy(),
		"intersight_hyperflex_ucsm_config_policy":                            resourceHyperflexUcsmConfigPolicy(),
		"intersight_hyperflex_vcenter_config_policy":                         resourceHyperflexVcenterConfigPolicy(),
		"intersight_hyperflex_vm_import_operation":                           resourceHyperflexVmImportOperation(),
		"intersight_hyperflex_vm_restore_operation":                          resourceHyperflexVmRestoreOperation(),
		"intersight_iam_account":                                             resourceIamAccount(),
		"intersight_iam_account_experience":                                  resourceIamAccountExperience(),
		"intersight_iam_api_key":                                             resourceIamApiKey(),
		"intersight_iam_app_registration":                                    resourceIamAppRegistration(),
		"intersight_iam_certificate":                                         resourceIamCertificate(),
		"intersight_iam_certificate_request":                                 resourceIamCertificateRequest(),
		"intersight_iam_end_point_user":                                      resourceIamEndPointUser(),
		"intersight_iam_end_point_user_policy":                               resourceIamEndPointUserPolicy(),
		"intersight_iam_end_point_user_role":                                 resourceIamEndPointUserRole(),
		"intersight_iam_idp":                                                 resourceIamIdp(),
		"intersight_iam_ip_access_management":                                resourceIamIpAccessManagement(),
		"intersight_iam_ip_address":                                          resourceIamIpAddress(),
		"intersight_iam_ldap_group":                                          resourceIamLdapGroup(),
		"intersight_iam_ldap_policy":                                         resourceIamLdapPolicy(),
		"intersight_iam_ldap_provider":                                       resourceIamLdapProvider(),
		"intersight_iam_permission":                                          resourceIamPermission(),
		"intersight_iam_private_key_spec":                                    resourceIamPrivateKeySpec(),
		"intersight_iam_qualifier":                                           resourceIamQualifier(),
		"intersight_iam_resource_roles":                                      resourceIamResourceRoles(),
		"intersight_iam_session_limits":                                      resourceIamSessionLimits(),
		"intersight_iam_trust_point":                                         resourceIamTrustPoint(),
		"intersight_iam_user":                                                resourceIamUser(),
		"intersight_iam_user_group":                                          resourceIamUserGroup(),
		"intersight_ipmioverlan_policy":                                      resourceIpmioverlanPolicy(),
		"intersight_ippool_pool":                                             resourceIppoolPool(),
		"intersight_iqnpool_pool":                                            resourceIqnpoolPool(),
		"intersight_kubernetes_aci_cni_apic":                                 resourceKubernetesAciCniApic(),
		"intersight_kubernetes_aci_cni_profile":                              resourceKubernetesAciCniProfile(),
		"intersight_kubernetes_aci_cni_tenant_cluster_allocation":            resourceKubernetesAciCniTenantClusterAllocation(),
		"intersight_kubernetes_addon_definition":                             resourceKubernetesAddonDefinition(),
		"intersight_kubernetes_addon_policy":                                 resourceKubernetesAddonPolicy(),
		"intersight_kubernetes_addon_repository":                             resourceKubernetesAddonRepository(),
		"intersight_kubernetes_baremetal_node_profile":                       resourceKubernetesBaremetalNodeProfile(),
		"intersight_kubernetes_cluster":                                      resourceKubernetesCluster(),
		"intersight_kubernetes_cluster_addon_profile":                        resourceKubernetesClusterAddonProfile(),
		"intersight_kubernetes_cluster_profile":                              resourceKubernetesClusterProfile(),
		"intersight_kubernetes_container_runtime_policy":                     resourceKubernetesContainerRuntimePolicy(),
		"intersight_kubernetes_network_policy":                               resourceKubernetesNetworkPolicy(),
		"intersight_kubernetes_node_group_profile":                           resourceKubernetesNodeGroupProfile(),
		"intersight_kubernetes_sys_config_policy":                            resourceKubernetesSysConfigPolicy(),
		"intersight_kubernetes_trusted_registries_policy":                    resourceKubernetesTrustedRegistriesPolicy(),
		"intersight_kubernetes_version":                                      resourceKubernetesVersion(),
		"intersight_kubernetes_version_policy":                               resourceKubernetesVersionPolicy(),
		"intersight_kubernetes_virtual_machine_infra_config_policy":          resourceKubernetesVirtualMachineInfraConfigPolicy(),
		"intersight_kubernetes_virtual_machine_infrastructure_provider":      resourceKubernetesVirtualMachineInfrastructureProvider(),
		"intersight_kubernetes_virtual_machine_instance_type":                resourceKubernetesVirtualMachineInstanceType(),
		"intersight_kubernetes_virtual_machine_node_profile":                 resourceKubernetesVirtualMachineNodeProfile(),
		"intersight_kvm_policy":                                              resourceKvmPolicy(),
		"intersight_kvm_session":                                             resourceKvmSession(),
		"intersight_kvm_tunnel":                                              resourceKvmTunnel(),
		"intersight_license_iks_license_count":                               resourceLicenseIksLicenseCount(),
		"intersight_license_iwo_license_count":                               resourceLicenseIwoLicenseCount(),
		"intersight_license_license_info":                                    resourceLicenseLicenseInfo(),
		"intersight_license_license_reservation_op":                          resourceLicenseLicenseReservationOp(),
		"intersight_macpool_pool":                                            resourceMacpoolPool(),
		"intersight_memory_persistent_memory_policy":                         resourceMemoryPersistentMemoryPolicy(),
		"intersight_networkconfig_policy":                                    resourceNetworkconfigPolicy(),
		"intersight_notification_account_subscription":                       resourceNotificationAccountSubscription(),
		"intersight_ntp_policy":                                              resourceNtpPolicy(),
		"intersight_oauth_authorization":                                     resourceOauthAuthorization(),
		"intersight_oprs_deployment":                                         resourceOprsDeployment(),
		"intersight_oprs_sync_target_list_message":                           resourceOprsSyncTargetListMessage(),
		"intersight_organization_organization":                               resourceOrganizationOrganization(),
		"intersight_os_bulk_install_info":                                    resourceOsBulkInstallInfo(),
		"intersight_os_configuration_file":                                   resourceOsConfigurationFile(),
		"intersight_os_install":                                              resourceOsInstall(),
		"intersight_power_policy":                                            resourcePowerPolicy(),
		"intersight_recovery_backup_config_policy":                           resourceRecoveryBackupConfigPolicy(),
		"intersight_recovery_backup_profile":                                 resourceRecoveryBackupProfile(),
		"intersight_recovery_on_demand_backup":                               resourceRecoveryOnDemandBackup(),
		"intersight_recovery_restore":                                        resourceRecoveryRestore(),
		"intersight_recovery_schedule_config_policy":                         resourceRecoveryScheduleConfigPolicy(),
		"intersight_resource_group":                                          resourceResourceGroup(),
		"intersight_resource_reservation":                                    resourceResourceReservation(),
		"intersight_resourcepool_pool":                                       resourceResourcepoolPool(),
		"intersight_sdcard_policy":                                           resourceSdcardPolicy(),
		"intersight_sdwan_profile":                                           resourceSdwanProfile(),
		"intersight_sdwan_router_node":                                       resourceSdwanRouterNode(),
		"intersight_sdwan_router_policy":                                     resourceSdwanRouterPolicy(),
		"intersight_sdwan_vmanage_account_policy":                            resourceSdwanVmanageAccountPolicy(),
		"intersight_server_config_import":                                    resourceServerConfigImport(),
		"intersight_server_profile":                                          resourceServerProfile(),
		"intersight_server_profile_template":                                 resourceServerProfileTemplate(),
		"intersight_smtp_policy":                                             resourceSmtpPolicy(),
		"intersight_snmp_policy":                                             resourceSnmpPolicy(),
		"intersight_software_appliance_distributable":                        resourceSoftwareApplianceDistributable(),
		"intersight_software_hcl_meta":                                       resourceSoftwareHclMeta(),
		"intersight_software_hyperflex_bundle_distributable":                 resourceSoftwareHyperflexBundleDistributable(),
		"intersight_software_hyperflex_distributable":                        resourceSoftwareHyperflexDistributable(),
		"intersight_software_release_meta":                                   resourceSoftwareReleaseMeta(),
		"intersight_software_solution_distributable":                         resourceSoftwareSolutionDistributable(),
		"intersight_software_ucsd_bundle_distributable":                      resourceSoftwareUcsdBundleDistributable(),
		"intersight_software_ucsd_distributable":                             resourceSoftwareUcsdDistributable(),
		"intersight_softwarerepository_authorization":                        resourceSoftwarerepositoryAuthorization(),
		"intersight_softwarerepository_category_mapper":                      resourceSoftwarerepositoryCategoryMapper(),
		"intersight_softwarerepository_category_mapper_model":                resourceSoftwarerepositoryCategoryMapperModel(),
		"intersight_softwarerepository_category_support_constraint":          resourceSoftwarerepositoryCategorySupportConstraint(),
		"intersight_softwarerepository_operating_system_file":                resourceSoftwarerepositoryOperatingSystemFile(),
		"intersight_softwarerepository_release":                              resourceSoftwarerepositoryRelease(),
		"intersight_sol_policy":                                              resourceSolPolicy(),
		"intersight_ssh_policy":                                              resourceSshPolicy(),
		"intersight_storage_drive_group":                                     resourceStorageDriveGroup(),
		"intersight_storage_storage_policy":                                  resourceStorageStoragePolicy(),
		"intersight_syslog_policy":                                           resourceSyslogPolicy(),
		"intersight_tam_advisory_count":                                      resourceTamAdvisoryCount(),
		"intersight_tam_advisory_definition":                                 resourceTamAdvisoryDefinition(),
		"intersight_tam_advisory_info":                                       resourceTamAdvisoryInfo(),
		"intersight_tam_advisory_instance":                                   resourceTamAdvisoryInstance(),
		"intersight_tam_security_advisory":                                   resourceTamSecurityAdvisory(),
		"intersight_techsupportmanagement_collection_control_policy":         resourceTechsupportmanagementCollectionControlPolicy(),
		"intersight_techsupportmanagement_tech_support_bundle":               resourceTechsupportmanagementTechSupportBundle(),
		"intersight_terraform_executor":                                      resourceTerraformExecutor(),
		"intersight_thermal_policy":                                          resourceThermalPolicy(),
		"intersight_uuidpool_pool":                                           resourceUuidpoolPool(),
		"intersight_virtualization_cisco_hypervisor_manager":                 resourceVirtualizationCiscoHypervisorManager(),
		"intersight_virtualization_esxi_console":                             resourceVirtualizationEsxiConsole(),
		"intersight_virtualization_iwe_datacenter":                           resourceVirtualizationIweDatacenter(),
		"intersight_virtualization_virtual_disk":                             resourceVirtualizationVirtualDisk(),
		"intersight_virtualization_virtual_machine":                          resourceVirtualizationVirtualMachine(),
		"intersight_virtualization_virtual_network":                          resourceVirtualizationVirtualNetwork(),
		"intersight_vmedia_policy":                                           resourceVmediaPolicy(),
		"intersight_vmrc_console":                                            resourceVmrcConsole(),
		"intersight_vnc_console":                                             resourceVncConsole(),
		"intersight_vnic_eth_adapter_policy":                                 resourceVnicEthAdapterPolicy(),
		"intersight_vnic_eth_if":                                             resourceVnicEthIf(),
		"intersight_vnic_eth_network_policy":                                 resourceVnicEthNetworkPolicy(),
		"intersight_vnic_eth_qos_policy":                                     resourceVnicEthQosPolicy(),
		"intersight_vnic_fc_adapter_policy":                                  resourceVnicFcAdapterPolicy(),
		"intersight_vnic_fc_if":                                              resourceVnicFcIf(),
		"intersight_vnic_fc_network_policy":                                  resourceVnicFcNetworkPolicy(),
		"intersight_vnic_fc_qos_policy":                                      resourceVnicFcQosPolicy(),
		"intersight_vnic_iscsi_adapter_policy":                               resourceVnicIscsiAdapterPolicy(),
		"intersight_vnic_iscsi_boot_policy":                                  resourceVnicIscsiBootPolicy(),
		"intersight_vnic_iscsi_static_target_policy":                         resourceVnicIscsiStaticTargetPolicy(),
		"intersight_vnic_lan_connectivity_policy":                            resourceVnicLanConnectivityPolicy(),
		"intersight_vnic_san_connectivity_policy":                            resourceVnicSanConnectivityPolicy(),
		"intersight_vrf_vrf":                                                 resourceVrfVrf(),
		"intersight_workflow_ansible_batch_executor":                         resourceWorkflowAnsibleBatchExecutor(),
		"intersight_workflow_batch_api_executor":                             resourceWorkflowBatchApiExecutor(),
		"intersight_workflow_custom_data_type_definition":                    resourceWorkflowCustomDataTypeDefinition(),
		"intersight_workflow_error_response_handler":                         resourceWorkflowErrorResponseHandler(),
		"intersight_workflow_rollback_workflow":                              resourceWorkflowRollbackWorkflow(),
		"intersight_workflow_solution_action_definition":                     resourceWorkflowSolutionActionDefinition(),
		"intersight_workflow_solution_action_instance":                       resourceWorkflowSolutionActionInstance(),
		"intersight_workflow_solution_definition":                            resourceWorkflowSolutionDefinition(),
		"intersight_workflow_solution_instance":                              resourceWorkflowSolutionInstance(),
		"intersight_workflow_solution_output":                                resourceWorkflowSolutionOutput(),
		"intersight_workflow_ssh_batch_executor":                             resourceWorkflowSshBatchExecutor(),
		"intersight_workflow_task_definition":                                resourceWorkflowTaskDefinition(),
		"intersight_workflow_workflow_definition":                            resourceWorkflowWorkflowDefinition(),
		"intersight_workflow_workflow_info":                                  resourceWorkflowWorkflowInfo(),
	}
}
