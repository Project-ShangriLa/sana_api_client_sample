# Sana API Ruby Sample


## follower_status.rb

指定した年/クールのTwitterアカウントのフォロワー数リストを表示

```
 bundle exe ruby follower_status.rb 2015 4
```

結果

```
     ketsuekigatakun 5659
         anime_lance 8685
            DwD_info 22470
          g_tekketsu 35960
      DDhokuto_anime 2572
      owarino_seraph 95786
      anime_kagewani 1625
              te_kyu 12599
         rainy_cocoa 4224
        hokuto_15aji 11417
       fafnerproject 29933
     Anime_W_Trigger 41316
          onsenyosei 2685
         usagi_anime 134950
            ho_anime 7633
         f_noitamina 12723
     nisioisin_anime 131155
            hstar_mu 26364
        cometlucifer 8161
     anime_kyojinchu 30973
        SakurakoKujo 10738
     UtawareOfficial 14501
         shinmaimaou 14722
     animehaikyu_com 286058
      anime_yuruyuri 92904
           ittoshura 15200
           anime_k11 92573
           opm_anime 52513
     animekindaichiR 5132
        ariaAA_anime 11440
        asterisk_war 15870
        anime_syomin 11946
      35shoutai_info 7946
        anime_somera 2553
     Hackadoll_anime 7120
        lupinIII_4th 20991
      anime_dialover 68508
         conrevoinfo 8273
           anime_ybj 9113
       valkyriedrive 6383
        komori_anime 2570
         noragami_PR 81402
         osomatsu_PR 129156
```


## follower_history.rb

指定したアニメ公式Twitterアカウントのフォロワー履歴を取得

```
 bundle exe ruby follower_history.rb usagi_anime
```

or 

```
 bundle exe ruby follower_history.rb usagi_anime 1446901208
```

結果

```
134964 2015-11-10 00:30:07 1447083007
134942 2015-11-10 00:00:07 1447081207
134924 2015-11-09 23:30:06 1447079406
134900 2015-11-09 23:00:07 1447077607
134882 2015-11-09 22:30:07 1447075807
134868 2015-11-09 22:00:07 1447074007
134851 2015-11-09 21:30:06 1447072206
134838 2015-11-09 21:00:07 1447070407
134824 2015-11-09 20:30:06 1447068606
134809 2015-11-09 20:00:06 1447066806
134791 2015-11-09 19:30:07 1447065007
134766 2015-11-09 19:00:06 1447063206
134759 2015-11-09 18:30:07 1447061407
134755 2015-11-09 18:00:07 1447059607
134747 2015-11-09 17:30:06 1447057806
134745 2015-11-09 17:00:06 1447056006
134741 2015-11-09 16:30:06 1447054206
134734 2015-11-09 16:00:07 1447052407
134730 2015-11-09 15:30:07 1447050607
134724 2015-11-09 15:00:06 1447048806
134720 2015-11-09 14:30:06 1447047006
134718 2015-11-09 14:00:07 1447045207
134718 2015-11-09 13:30:06 1447043406
134715 2015-11-09 13:00:06 1447041606
134711 2015-11-09 12:30:06 1447039806
134712 2015-11-09 12:00:06 1447038006
```