read -t 5 -p "Migrate UP or DOWN: " option
cd ~/boot_projects/blog_aggregator/sql/schema;
goose postgres postgres://user:pass@localhost:5432/blog_db  $option