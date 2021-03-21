require 'shangrila'

account = ARGV[0]
days = ARGV[1]
api_host = ARGV[2]

sana = ARGV[2] == nil ? Shangrila::Sana.new() : Shangrila::Sana.new(api_host)

#指定したアカウントのフォロワー履歴を取得
history_list = sana.follower_history_daily(account, days)

history_list.each do |history|
 puts sprintf('%s,%d', Time.at(history['updated_at']).strftime('%Y-%m-%d'),history['follower'] )
end