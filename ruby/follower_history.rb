require 'shangrila'

account = ARGV[0]
end_date_unixtimestamp = ARGV[1]
api_host = ARGV[2]

#指定したアカウントのフォロワー履歴を取得
history_list = Shangrila::Sana.new(api_host).follower_history(account, end_date_unixtimestamp)

history_list.each do |history|
 puts sprintf('%d %s %d', history['follower'], Time.at(history['updated_at']).strftime('%Y-%m-%d %H:%M:%S'), history['updated_at'] )
end