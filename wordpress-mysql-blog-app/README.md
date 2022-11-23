## WordPress-MySQL blog application on Red Hat OpenShift Container Platform v3.9

[![Docker Repository on Quay](https://quay.io/repository/alberthua/wordpress/status "Docker Repository on Quay")](https://quay.io/repository/alberthua/wordpress) [![Docker Repository on Quay](https://quay.io/repository/alberthua/mysql-55-rhel7/status "Docker Repository on Quay")](https://quay.io/repository/alberthua/mysql-55-rhel7)

- Practice environment: Red Hat OpenShift Container Platform v3.9

- Deploy blog application:

  - Create myblog project.

  - Create static pv and pvc for mysql in myblog project.

  - Create mysql pod and service as database backend where not use statefulset resource.

  - Use the same method to create static pv and pvc for wordpress.

  - ✨ Create serviceaccount named `wp-admin` and associated `anyuid` scc to wp-admin.

    > Note: 
    > Because wordpress image from docker.io in which root run the container, but in openshift use random user to run pod instead of root.

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
