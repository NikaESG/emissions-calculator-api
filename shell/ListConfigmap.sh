ns=$1

# These Lines are adding files into operation paths.
echo $PATH
export PATH="/Users/user/.rd/bin:$PATH"
export PATH="/usr/local/bin/:$PATH"
echo $PATH

kubectl get configmap --namespace $ns