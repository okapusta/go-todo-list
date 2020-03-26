require 'sequel'
require 'faker'

DB = Sequel.connect(ENV['DATABASE_URL'] || 'mysql2://root:@localhost/todos_development')

todos = DB[:todos]

puts "Seeding todos table..."
10.times do
  content = Faker::Company.bs
  puts "Creating todo #{content}"
  todos.insert(content: content)
end
puts "Done."
