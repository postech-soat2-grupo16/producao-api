provider "aws" {
  region = var.aws_region
}

#Configuração do Terraform State
terraform {
  backend "s3" {
    bucket = "terraform-state-soat"
    key    = "producao-api/terraform.tfstate"
    region = "us-east-1"

    dynamodb_table = "terraform-state-soat-locking"
    encrypt        = true
  }
}

### Target Group + Listener

resource "aws_lb_target_group" "tg_producao_api" {
  name        = "target-group-producao-api"
  port        = 8000
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = var.vpc_id

  health_check {
    enabled             = true
    interval            = 30
    matcher             = "200-299"
    path                = "/producao_pedidos"
    port                = "traffic-port"
    protocol            = "HTTP"
    timeout             = 5
    healthy_threshold   = 5
    unhealthy_threshold = 2
  }

  tags = {
    infra   = "target-group-producao-api"
    service = "producao"
  }
}

# Listener Rule that forwards the request to producao-api TG
resource "aws_lb_listener_rule" "listener_producao_api" {
  listener_arn = var.alb_fastfood_listener_arn
  priority     = 300

  condition {
    path_pattern {
      values = ["/producao_pedidos*"]
    }
  }

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.tg_producao_api.arn
  }

  tags = {
    Name    = "alb-listener-producao"
    infra   = "alb-listener-producao"
    service = "producao"
  }
}


### Task Config ###
resource "aws_ecs_task_definition" "task_definition_producao_api" {
  family                   = "task-definition-producao-api"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = var.execution_role_ecs
  task_role_arn            = var.execution_role_ecs

  cpu    = 512
  memory = 1024

  container_definitions = jsonencode([
    {
      name      = "container-producao-api"
      image     = var.ecr_image
      cpu       = 512,
      memory    = 1024,
      essential = true,
      portMappings = [
        {
          containerPort = 8000
          hostPort      = 8000
          protocol      = "tcp"
          appProtocol   = "http"
        }
      ]

      environment = [
        { "name" : "DATABASE_URL", "value" : var.db_url },
        { "name" : "QUEUE_URL", "value" : var.sqs_url },
        { "name" : "IS_LOCAL", "value" : "false" }
      ]


      logConfiguration = {
        logDriver = "awslogs",
        options = {
          awslogs-create-group  = "true",
          awslogs-group         = "producao-api-container",
          awslogs-region        = "us-east-1",
          awslogs-stream-prefix = "producao"
        }
      },
    }
  ])

  runtime_platform {
    operating_system_family = "LINUX"
    cpu_architecture        = "X86_64"
  }

  tags = {
    infra    = "task-definition-producao"
    services = "producao"
  }
}

resource "aws_ecs_service" "ecs_service_producao_api" {
  name                              = "ecs-service-producao-api"
  cluster                           = var.ecs_cluster
  task_definition                   = aws_ecs_task_definition.task_definition_producao_api.id
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  desired_count                     = var.desired_tasks
  health_check_grace_period_seconds = 30

  network_configuration {
    subnets = [
      var.subnet_a,
      var.subnet_b
    ]
    security_groups  = [var.sg_cluster_ecs]
    assign_public_ip = false
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.tg_producao_api.arn
    container_name   = "container-producao-api"
    container_port   = 8000
  }

  tags = {
    infra    = "ecs-service-producao"
    services = "producao"
  }
}