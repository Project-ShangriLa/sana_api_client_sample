require 'shangrila'

account = ARGV[0]
days = ARGV[1]
api_host = ARGV[2]

#指定したアカウントのフォロワー履歴を取得
history_list = Shangrila::Sana.new(api_host).follower_history_daily(account, days)

history_list.each do |history|
 puts sprintf('%s,%d', Time.at(history['updated_at']).strftime('%Y-%m-%d'),history['follower'] )
end