{
        "changed": false,
        "msg": "All items completed",
        "results": [
            {
                "ansible_loop_var": "item",
                "changed": false,
                "changes": [],
                "failed": false,
                "instance_ids": [
                    "i-013fdb1004582f17f"
                ],
                "instances": [
                    {
                        "ami_launch_index": 0,
                        "architecture": "x86_64",
                        "block_device_mappings": [
                            {
                                "device_name": "/dev/sda1",
                                "ebs": {
                                    "attach_time": "2025-04-07T13:08:09+00:00",
                                    "delete_on_termination": true,
                                    "status": "attached",
                                    "volume_id": "vol-0dbcd2242367ba78a"
                                }
                            }
                        ],
                        "boot_mode": "uefi-preferred",
                        "capacity_reservation_specification": {
                            "capacity_reservation_preference": "open"
                        },
                        "client_token": "1105f114459647e5a0b188b6e6d512d5",
                        "cpu_options": {
                            "core_count": 1,
                            "threads_per_core": 2
                        },
                        "current_instance_boot_mode": "uefi",
                        "ebs_optimized": false,
                        "ena_support": true,
                        "enclave_options": {
                            "enabled": false
                        },
                        "hibernation_options": {
                            "configured": false
                        },
                        "hypervisor": "xen",
                        "image_id": "ami-09c813fb71547fc4f",
                        "instance_id": "i-013fdb1004582f17f",
                        "instance_type": "t3.micro",
                        "launch_time": "2025-04-07T13:08:09+00:00",
                        "maintenance_options": {
                            "auto_recovery": "default"
                        },
                        "metadata_options": {
                            "http_endpoint": "enabled",
                            "http_protocol_ipv6": "disabled",
                            "http_put_response_hop_limit": 1,
                            "http_tokens": "optional",
                            "instance_metadata_tags": "disabled",
                            "state": "applied"
                        },
                        "monitoring": {
                            "state": "disabled"
                        },
                        "network_interfaces": [
                            {
                                "association": {
                                    "ip_owner_id": "amazon",
                                    "public_dns_name": "ec2-44-200-50-254.compute-1.amazonaws.com",
                                    "public_ip": "44.200.50.254"
                                },
                                "attachment": {
                                    "attach_time": "2025-04-07T13:08:09+00:00",
                                    "attachment_id": "eni-attach-00bbb0ecfadc39502",
                                    "delete_on_termination": true,
                                    "device_index": 0,
                                    "network_card_index": 0,
                                    "status": "attached"
                                },
                                "description": "",
                                "groups": [
                                    {
                                        "group_id": "sg-08f49690bf42656e1",
                                        "group_name": "allow-all"
                                    }
                                ],
                                "interface_type": "interface",
                                "ipv6_addresses": [],
                                "mac_address": "02:b8:ac:28:f1:97",
                                "network_interface_id": "eni-00debe96fa49c8e4b",
                                "operator": {
                                    "managed": false
                                },
                                "owner_id": "314146304089",
                                "private_dns_name": "ip-172-31-15-85.ec2.internal",
                                "private_ip_address": "172.31.15.85",
                                "private_ip_addresses": [
                                    {
                                        "association": {
                                            "ip_owner_id": "amazon",
                                            "public_dns_name": "ec2-44-200-50-254.compute-1.amazonaws.com",
                                            "public_ip": "44.200.50.254"
                                        },
                                        "primary": true,
                                        "private_dns_name": "ip-172-31-15-85.ec2.internal",
                                        "private_ip_address": "172.31.15.85"
                                    }
                                ],
                                "source_dest_check": true,
                                "status": "in-use",
                                "subnet_id": "subnet-0e40ff3b0a56e2315",
                                "vpc_id": "vpc-00716bedb7bc0c3d8"
                            }
                        ],
                        "network_performance_options": {
                            "bandwidth_weighting": "default"
                        },
                        "operator": {
                            "managed": false
                        },
                        "placement": {
                            "availability_zone": "us-east-1c",
                            "group_name": "",
                            "tenancy": "default"
                        },
                        "platform_details": "Red Hat Enterprise Linux",
                        "private_dns_name": "ip-172-31-15-85.ec2.internal",
                        "private_dns_name_options": {
                            "enable_resource_name_dns_a_record": false,
                            "enable_resource_name_dns_aaaa_record": false,
                            "hostname_type": "ip-name"
                        },
                        "private_ip_address": "172.31.15.85",
                        "product_codes": [],
                        "public_dns_name": "ec2-44-200-50-254.compute-1.amazonaws.com",
                        "public_ip_address": "44.200.50.254",
                        "root_device_name": "/dev/sda1",
                        "root_device_type": "ebs",
                        "security_groups": [
                            {
                                "group_id": "sg-08f49690bf42656e1",
                                "group_name": "allow-all"
                            }
                        ],
                        "source_dest_check": true,
                        "state": {
                            "code": 16,
                            "name": "running"
                        },
                        "state_transition_reason": "",
                        "subnet_id": "subnet-0e40ff3b0a56e2315",
                        "tags": {
                            "Name": "mysql"
                        },
                        "usage_operation": "RunInstances:0010",
                        "usage_operation_update_time": "2025-04-07T13:08:09+00:00",
                        "virtualization_type": "hvm",
                        "vpc_id": "vpc-00716bedb7bc0c3d8"
                    }
                ],
                "invocation": {
                    "module_args": {
                        "aap_callback": null,
                        "access_key": null,
                        "availability_zone": null,
                        "aws_ca_bundle": "/etc/pki/tls/certs/ca-bundle.crt",
                        "aws_config": null,
                        "count": null,
                        "cpu_credit_specification": null,
                        "cpu_options": null,
                        "debug_botocore_endpoint_logs": false,
                        "detailed_monitoring": null,
                        "ebs_optimized": null,
                        "endpoint_url": null,
                        "exact_count": null,
                        "filters": {
                            "image-id": [
                                "ami-09c813fb71547fc4f"
                            ],
                            "instance-state-name": [
                                "pending",
                                "running",
                                "stopping",
                                "stopped"
                            ],
                            "subnet-id": [
                                "subnet-0e40ff3b0a56e2315"
                            ],
                            "tag:Name": [
                                "mysql"
                            ]
                        },
                        "hibernation_options": false,
                        "iam_instance_profile": null,
                        "image": null,
                        "image_id": "ami-09c813fb71547fc4f",
                        "instance_ids": [],
                        "instance_initiated_shutdown_behavior": null,
                        "instance_type": "t3.micro",
                        "key_name": null,
                        "launch_template": null,
                        "metadata_options": null,
                        "name": "mysql",
                        "network": null,
                        "placement_group": null,
                        "profile": null,
                        "purge_tags": true,
                        "region": "us-east-1",
                        "secret_key": null,
                        "security_group": "sg-08f49690bf42656e1",
                        "security_groups": [],
                        "session_token": null,
                        "state": "present",
                        "subnet_id": "subnet-0e40ff3b0a56e2315",
                        "tags": null,
                        "tenancy": null,
                        "termination_protection": null,
                        "user_data": null,
                        "validate_certs": true,
                        "volumes": null,
                        "vpc_subnet_id": "subnet-0e40ff3b0a56e2315",
                        "wait": true,
                        "wait_timeout": 600
                    }
                },
                "item": "mysql"
            },
            {
                "ansible_loop_var": "item",
                "changed": false,
                "changes": [],
                "failed": false,
                "instance_ids": [
                    "i-06a763e01534cf331"
                ],
                "instances": [
                    {
                        "ami_launch_index": 0,
                        "architecture": "x86_64",
                        "block_device_mappings": [
                            {
                                "device_name": "/dev/sda1",
                                "ebs": {
                                    "attach_time": "2025-04-07T13:08:13+00:00",
                                    "delete_on_termination": true,
                                    "status": "attached",
                                    "volume_id": "vol-0e6199ed3c9f705d1"
                                }
                            }
                        ],
                        "boot_mode": "uefi-preferred",
                        "capacity_reservation_specification": {
                            "capacity_reservation_preference": "open"
                        },
                        "client_token": "cdc9824439f043cf91407b7a1f5a36ce",
                        "cpu_options": {
                            "core_count": 1,
                            "threads_per_core": 2
                        },
                        "current_instance_boot_mode": "uefi",
                        "ebs_optimized": false,
                        "ena_support": true,
                        "enclave_options": {
                            "enabled": false
                        },
                        "hibernation_options": {
                            "configured": false
                        },
                        "hypervisor": "xen",
                        "image_id": "ami-09c813fb71547fc4f",
                        "instance_id": "i-06a763e01534cf331",
                        "instance_type": "t3.micro",
                        "launch_time": "2025-04-07T13:08:13+00:00",
                        "maintenance_options": {
                            "auto_recovery": "default"
                        },
                        "metadata_options": {
                            "http_endpoint": "enabled",
                            "http_protocol_ipv6": "disabled",
                            "http_put_response_hop_limit": 1,
                            "http_tokens": "optional",
                            "instance_metadata_tags": "disabled",
                            "state": "applied"
                        },
                        "monitoring": {
                            "state": "disabled"
                        },
                        "network_interfaces": [
                            {
                                "association": {
                                    "ip_owner_id": "amazon",
                                    "public_dns_name": "ec2-34-234-236-173.compute-1.amazonaws.com",
                                    "public_ip": "34.234.236.173"
                                },
                                "attachment": {
                                    "attach_time": "2025-04-07T13:08:13+00:00",
                                    "attachment_id": "eni-attach-0e80f331ee8957542",
                                    "delete_on_termination": true,
                                    "device_index": 0,
                                    "network_card_index": 0,
                                    "status": "attached"
                                },
                                "description": "",
                                "groups": [
                                    {
                                        "group_id": "sg-08f49690bf42656e1",
                                        "group_name": "allow-all"
                                    }
                                ],
                                "interface_type": "interface",
                                "ipv6_addresses": [],
                                "mac_address": "02:83:8a:21:b7:69",
                                "network_interface_id": "eni-0553b09cd6e5071ec",
                                "operator": {
                                    "managed": false
                                },
                                "owner_id": "314146304089",
                                "private_dns_name": "ip-172-31-11-122.ec2.internal",
                                "private_ip_address": "172.31.11.122",
                                "private_ip_addresses": [
                                    {
                                        "association": {
                                            "ip_owner_id": "amazon",
                                            "public_dns_name": "ec2-34-234-236-173.compute-1.amazonaws.com",
                                            "public_ip": "34.234.236.173"
                                        },
                                        "primary": true,
                                        "private_dns_name": "ip-172-31-11-122.ec2.internal",
                                        "private_ip_address": "172.31.11.122"
                                    }
                                ],
                                "source_dest_check": true,
                                "status": "in-use",
                                "subnet_id": "subnet-0e40ff3b0a56e2315",
                                "vpc_id": "vpc-00716bedb7bc0c3d8"
                            }
                        ],
                        "network_performance_options": {
                            "bandwidth_weighting": "default"
                        },
                        "operator": {
                            "managed": false
                        },
                        "placement": {
                            "availability_zone": "us-east-1c",
                            "group_name": "",
                            "tenancy": "default"
                        },
                        "platform_details": "Red Hat Enterprise Linux",
                        "private_dns_name": "ip-172-31-11-122.ec2.internal",
                        "private_dns_name_options": {
                            "enable_resource_name_dns_a_record": false,
                            "enable_resource_name_dns_aaaa_record": false,
                            "hostname_type": "ip-name"
                        },
                        "private_ip_address": "172.31.11.122",
                        "product_codes": [],
                        "public_dns_name": "ec2-34-234-236-173.compute-1.amazonaws.com",
                        "public_ip_address": "34.234.236.173",
                        "root_device_name": "/dev/sda1",
                        "root_device_type": "ebs",
                        "security_groups": [
                            {
                                "group_id": "sg-08f49690bf42656e1",
                                "group_name": "allow-all"
                            }
                        ],
                        "source_dest_check": true,
                        "state": {
                            "code": 16,
                            "name": "running"
                        },
                        "state_transition_reason": "",
                        "subnet_id": "subnet-0e40ff3b0a56e2315",
                        "tags": {
                            "Name": "backend"
                        },
                        "usage_operation": "RunInstances:0010",
                        "usage_operation_update_time": "2025-04-07T13:08:13+00:00",
                        "virtualization_type": "hvm",
                        "vpc_id": "vpc-00716bedb7bc0c3d8"
                    }
                ],
                "invocation": {
                    "module_args": {
                        "aap_callback": null,
                        "access_key": null,
                        "availability_zone": null,
                        "aws_ca_bundle": "/etc/pki/tls/certs/ca-bundle.crt",
                        "aws_config": null,
                        "count": null,
                        "cpu_credit_specification": null,
                        "cpu_options": null,
                        "debug_botocore_endpoint_logs": false,
                        "detailed_monitoring": null,
                        "ebs_optimized": null,
                        "endpoint_url": null,
                        "exact_count": null,
                        "filters": {
                            "image-id": [
                                "ami-09c813fb71547fc4f"
                            ],
                            "instance-state-name": [
                                "pending",
                                "running",
                                "stopping",
                                "stopped"
                            ],
                            "subnet-id": [
                                "subnet-0e40ff3b0a56e2315"
                            ],
                            "tag:Name": [
                                "backend"
                            ]
                        },
                        "hibernation_options": false,
                        "iam_instance_profile": null,
                        "image": null,
                        "image_id": "ami-09c813fb71547fc4f",
                        "instance_ids": [],
                        "instance_initiated_shutdown_behavior": null,
                        "instance_type": "t3.micro",
                        "key_name": null,
                        "launch_template": null,
                        "metadata_options": null,
                        "name": "backend",
                        "network": null,
                        "placement_group": null,
                        "profile": null,
                        "purge_tags": true,
                        "region": "us-east-1",
                        "secret_key": null,
                        "security_group": "sg-08f49690bf42656e1",
                        "security_groups": [],
                        "session_token": null,
                        "state": "present",
                        "subnet_id": "subnet-0e40ff3b0a56e2315",
                        "tags": null,
                        "tenancy": null,
                        "termination_protection": null,
                        "user_data": null,
                        "validate_certs": true,
                        "volumes": null,
                        "vpc_subnet_id": "subnet-0e40ff3b0a56e2315",
                        "wait": true,
                        "wait_timeout": 600
                    }
                },
                "item": "backend"
            },
            {
                "ansible_loop_var": "item",
                "changed": false,
                "changes": [],
                "failed": false,
                "instance_ids": [
                    "i-0e33079622161adf3"
                ],
                "instances": [
                    {
                        "ami_launch_index": 0,
                        "architecture": "x86_64",
                        "block_device_mappings": [
                            {
                                "device_name": "/dev/sda1",
                                "ebs": {
                                    "attach_time": "2025-04-07T13:08:17+00:00",
                                    "delete_on_termination": true,
                                    "status": "attached",
                                    "volume_id": "vol-0f015b13215e2e5d4"
                                }
                            }
                        ],
                        "boot_mode": "uefi-preferred",
                        "capacity_reservation_specification": {
                            "capacity_reservation_preference": "open"
                        },
                        "client_token": "a55801b03e9b4a75872e39aa584ad738",
                        "cpu_options": {
                            "core_count": 1,
                            "threads_per_core": 2
                        },
                        "current_instance_boot_mode": "uefi",
                        "ebs_optimized": false,
                        "ena_support": true,
                        "enclave_options": {
                            "enabled": false
                        },
                        "hibernation_options": {
                            "configured": false
                        },
                        "hypervisor": "xen",
                        "image_id": "ami-09c813fb71547fc4f",
                        "instance_id": "i-0e33079622161adf3",
                        "instance_type": "t3.micro",
                        "launch_time": "2025-04-07T13:08:17+00:00",
                        "maintenance_options": {
                            "auto_recovery": "default"
                        },
                        "metadata_options": {
                            "http_endpoint": "enabled",
                            "http_protocol_ipv6": "disabled",
                            "http_put_response_hop_limit": 1,
                            "http_tokens": "optional",
                            "instance_metadata_tags": "disabled",
                            "state": "applied"
                        },
                        "monitoring": {
                            "state": "disabled"
                        },
                        "network_interfaces": [
                            {
                                "association": {
                                    "ip_owner_id": "amazon",
                                    "public_dns_name": "ec2-44-213-105-17.compute-1.amazonaws.com",
                                    "public_ip": "44.213.105.17"
                                },
                                "attachment": {
                                    "attach_time": "2025-04-07T13:08:17+00:00",
                                    "attachment_id": "eni-attach-00f01bd1f0191dfa2",
                                    "delete_on_termination": true,
                                    "device_index": 0,
                                    "network_card_index": 0,
                                    "status": "attached"
                                },
                                "description": "",
                                "groups": [
                                    {
                                        "group_id": "sg-08f49690bf42656e1",
                                        "group_name": "allow-all"
                                    }
                                ],
                                "interface_type": "interface",
                                "ipv6_addresses": [],
                                "mac_address": "02:7e:1a:9d:d5:c1",
                                "network_interface_id": "eni-08052db7cda7ce052",
                                "operator": {
                                    "managed": false
                                },
                                "owner_id": "314146304089",
                                "private_dns_name": "ip-172-31-1-65.ec2.internal",
                                "private_ip_address": "172.31.1.65",
                                "private_ip_addresses": [
                                    {
                                        "association": {
                                            "ip_owner_id": "amazon",
                                            "public_dns_name": "ec2-44-213-105-17.compute-1.amazonaws.com",
                                            "public_ip": "44.213.105.17"
                                        },
                                        "primary": true,
                                        "private_dns_name": "ip-172-31-1-65.ec2.internal",
                                        "private_ip_address": "172.31.1.65"
                                    }
                                ],
                                "source_dest_check": true,
                                "status": "in-use",
                                "subnet_id": "subnet-0e40ff3b0a56e2315",
                                "vpc_id": "vpc-00716bedb7bc0c3d8"
                            }
                        ],
                        "network_performance_options": {
                            "bandwidth_weighting": "default"
                        },
                        "operator": {
                            "managed": false
                        },
                        "placement": {
                            "availability_zone": "us-east-1c",
                            "group_name": "",
                            "tenancy": "default"
                        },
                        "platform_details": "Red Hat Enterprise Linux",
                        "private_dns_name": "ip-172-31-1-65.ec2.internal",
                        "private_dns_name_options": {
                            "enable_resource_name_dns_a_record": false,
                            "enable_resource_name_dns_aaaa_record": false,
                            "hostname_type": "ip-name"
                        },
                        "private_ip_address": "172.31.1.65",
                        "product_codes": [],
                        "public_dns_name": "ec2-44-213-105-17.compute-1.amazonaws.com",
                        "public_ip_address": "44.213.105.17",
                        "root_device_name": "/dev/sda1",
                        "root_device_type": "ebs",
                        "security_groups": [
                            {
                                "group_id": "sg-08f49690bf42656e1",
                                "group_name": "allow-all"
                            }
                        ],
                        "source_dest_check": true,
                        "state": {
                            "code": 16,
                            "name": "running"
                        },
                        "state_transition_reason": "",
                        "subnet_id": "subnet-0e40ff3b0a56e2315",
                        "tags": {
                            "Name": "frontend"
                        },
                        "usage_operation": "RunInstances:0010",
                        "usage_operation_update_time": "2025-04-07T13:08:17+00:00",
                        "virtualization_type": "hvm",
                        "vpc_id": "vpc-00716bedb7bc0c3d8"
                    }
                ],
                "invocation": {
                    "module_args": {
                        "aap_callback": null,
                        "access_key": null,
                        "availability_zone": null,
                        "aws_ca_bundle": "/etc/pki/tls/certs/ca-bundle.crt",
                        "aws_config": null,
                        "count": null,
                        "cpu_credit_specification": null,
                        "cpu_options": null,
                        "debug_botocore_endpoint_logs": false,
                        "detailed_monitoring": null,
                        "ebs_optimized": null,
                        "endpoint_url": null,
                        "exact_count": null,
                        "filters": {
                            "image-id": [
                                "ami-09c813fb71547fc4f"
                            ],
                            "instance-state-name": [
                                "pending",
                                "running",
                                "stopping",
                                "stopped"
                            ],
                            "subnet-id": [
                                "subnet-0e40ff3b0a56e2315"
                            ],
                            "tag:Name": [
                                "frontend"
                            ]
                        },
                        "hibernation_options": false,
                        "iam_instance_profile": null,
                        "image": null,
                        "image_id": "ami-09c813fb71547fc4f",
                        "instance_ids": [],
                        "instance_initiated_shutdown_behavior": null,
                        "instance_type": "t3.micro",
                        "key_name": null,
                        "launch_template": null,
                        "metadata_options": null,
                        "name": "frontend",
                        "network": null,
                        "placement_group": null,
                        "profile": null,
                        "purge_tags": true,
                        "region": "us-east-1",
                        "secret_key": null,
                        "security_group": "sg-08f49690bf42656e1",
                        "security_groups": [],
                        "session_token": null,
                        "state": "present",
                        "subnet_id": "subnet-0e40ff3b0a56e2315",
                        "tags": null,
                        "tenancy": null,
                        "termination_protection": null,
                        "user_data": null,
                        "validate_certs": true,
                        "volumes": null,
                        "vpc_subnet_id": "subnet-0e40ff3b0a56e2315",
                        "wait": true,
                        "wait_timeout": 600
                    }
                },
                "item": "frontend"
            }
        ],
        "skipped": false
    }
