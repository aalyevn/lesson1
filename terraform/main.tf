variable "security_group_ids" {
  type    = list(string)
  default = ["sg-02762ba3be4dd8f04"]
}


resource "aws_elasticache_replication_group" "redis" {
  replication_group_id       = "redis707"
  description                = "Redis Replication Group"
  node_type                  = "cache.t2.micro"
  parameter_group_name       = "default.redis7.cluster.on"
  port                       = 6379
  automatic_failover_enabled = true
  security_group_ids         = var.security_group_ids
  subnet_group_name          = "redissubnet"
}
