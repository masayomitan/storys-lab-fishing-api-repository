CREATE TABLE tides (
    id INT AUTO_INCREMENT PRIMARY KEY ,
    prefecture_id INT NOT NULL COMMENT '都道府県のID',
    date DATE NOT NULL COMMENT '日付',
    sunrise_time TIME COMMENT '日の出時刻',
    sunset_time TIME COMMENT '日の入り時刻',
    moonrise_time TIME COMMENT '月の出時刻',
    moonset_time TIME COMMENT '月の入り時刻',
    weather VARCHAR(255) COMMENT '天気',
    temperature FLOAT COMMENT '気温',
    water_temperature FLOAT COMMENT '水温',
    wind_speed FLOAT COMMENT '風速',
    precipitation FLOAT COMMENT '降水量',
    wave_height FLOAT COMMENT '波の高さ',
    tide_flow VARCHAR(255) COMMENT '潮の流れ',
    low_tide_time_1 TIME COMMENT '第一低潮時刻',
    high_tide_time_1 TIME COMMENT '第一高潮時刻',
    low_tide_time_2 TIME COMMENT '第二低潮時刻',
    high_tide_time_2 TIME COMMENT '第二高潮時刻',
    low_tide_time_3 TIME COMMENT '第三低潮時刻',
    high_tide_time_3 TIME COMMENT '第三高潮時刻',
    tide_height_1 INT COMMENT '第一潮の高さ',
    tide_height_2 INT COMMENT '第二潮の高さ',
    tide_height_3 INT COMMENT '第三潮の高さ',
    moon_age FLOAT COMMENT '月齢',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時',
    FOREIGN KEY (prefecture_id) REFERENCES prefectures(id),
    INDEX idx_date (date),
    INDEX idx_weather (weather)
) COMMENT='潮汐に関連する情報を保持するテーブル';