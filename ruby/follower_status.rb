require 'shangrila'

year = ARGV[0]
cours = ARGV[1]
api_host = ARGV[2]

# 指定した年、クールのアニメ公式Twitterアカウントを取得(Sora[Master] API)
master =  Shangrila::Sora.new().get_flat_data(year, cours, ['twitter_account'])

twitter_account_list = master.flatten

# 指定したTwitterアカウントのフォロワー数を取得(Sana[Twitter] API)
sana = ARGV[2] == nil ? Shangrila::Sana.new() : Shangrila::Sana.new(api_host)
follower_status = sana.follower_status(twitter_account_list)

follower_status.each do |key,val|
  puts sprintf('%20s %d', key, val['follower'])
end
