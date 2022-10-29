## WordPress-MySQL Blog application on Red Hat OpenShift v3.9

- Deploy blog application:

  - Create myblog project.

  - Create pv and pvc for mysql in myblog project.

  - Create mysql pod and service as database backend where not use deploymentconfing and statefulset resource.

  - Use the same method to create pv and pvc for wordpress.

	- Create serviceaccount named wp-admin and associated anyuid scc to wp-admin.

		> Note: 
		> Because wordpress image from docker.io in which root run the container, but in openshift use random user to
    > run pod instead of root.

		```bash
		$ oc create serviceaccount wp-admin -n myblog
		$ oc adm policy add-scc-to-user anyuid -z wp-admin -n myblog
		```

  - Use deploymentconfig-service file to create wordpress frontend which discovery backend according to service name.

		```bash
		$ oc apply -f wordpress-deploymentconfig-service.yml -n myblog
		# wait for some minutes until wordpress running and ready
		```

  - Expose frontend service as route which can be access by client.

