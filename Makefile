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
		--image=weaveworks/flagger-loadtester:0.18.0 \
		--namespace staging \
		--context $(KUBE_CONTEXT) \
		--kubeconfig $(KUBECONFIG) \
		--command \
		-- /bin/sh -c "hey -z 1m -m GET -q 10 http://hello-gitops/"

