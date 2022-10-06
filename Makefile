KUBE_CONTEXT ?= default
KUBECONFIG ?= default

export KUBECONFIG KUBE_CONTEXT

check:
	which kubectl
	@echo "All pre-requisites available!"

load-test:
	@echo "In another window run:"
	@echo "watch kubectl get hpa --namespace dev --context $(KUBE_CONTEXT) --kubeconfig $(KUBECONFIG)"
	@echo ""
	@echo "Running load test.."
	kubectl run \
		-i --tty --rm --restart=Never \
		load-generator \
		--image=busybox:stable \
		--namespace dev \
		--context $(KUBE_CONTEXT) \
		--kubeconfig $(KUBECONFIG) \
		-- /bin/sh -c "while sleep 0.01; do wget -q -O/dev/null http://hello-gitops; done"

