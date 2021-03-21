require 'shangrila'

account = ARGV[0]
end_date_unixtimestamp = ARGV[1]
api_host = ARGV[2]

#指定したアカウントのフォロワー履歴を取得
sana = ARGV[2] == nil ? Shangrila::Sana.new() : Shangrila::Sana.new(api_host)
history_list = sana.follower_history(account, end_date_unixtimestamp)

history_list.each do |history|
 puts sprintf('%d %s %d', history['follower'], Time.at(history['updated_at']).strftime('%Y-%m-%d %H:%M:%S'), history['updated_at'] )
end