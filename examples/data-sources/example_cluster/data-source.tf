# Looks up a managed cluster by name.
data "example_cluster" "primary" {
  name = "Primary VA Cluster"
}

output "cluster_id" {
  value = data.example_cluster.primary.id
}
