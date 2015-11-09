require 'shangrila'

year = ARGV[0]
cours = ARGV[1]

# 指定した年、クールのアニメ公式Twitterアカウントを取得(Sora[Master] API)
master =  Shangrila::Sora.new().get_flat_data(year, cours, ['twitter_account'])

twitter_account_list = master.flatten

# 指定したTwitterアカウントのフォロワー数を取得(Sana[Twitter] API)
follower_status = Shangrila::Sana.new().follower_status(twitter_account_list)

follower_status.each do |key,val|
  puts sprintf('%20s %d', key, val['follower'])
end
