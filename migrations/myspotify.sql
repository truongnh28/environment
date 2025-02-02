-- MySQL dump 10.13  Distrib 8.0.31, for macos12.6 (x86_64)
--
-- Host: 127.0.0.1    Database: myspotify
-- ------------------------------------------------------
-- Server version	8.0.30

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `albums`
--

DROP TABLE IF EXISTS `albums`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `albums` (
                          `album_ID` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(40) NOT NULL,
                          `artist_ID` int NOT NULL,
                          `created_at` date DEFAULT NULL,
                          `upload_at` date DEFAULT NULL,
                          `cover_img` varchar(400) DEFAULT NULL,
                          PRIMARY KEY (`album_ID`),
                          KEY `artist_ID` (`artist_ID`),
                          CONSTRAINT `albums_ibfk_1` FOREIGN KEY (`artist_ID`) REFERENCES `artists` (`artist_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `albums`
--

LOCK TABLES `albums` WRITE;
/*!40000 ALTER TABLE `albums` DISABLE KEYS */;
INSERT INTO `albums` (`album_ID`, `name`, `artist_ID`, `created_at`, `upload_at`, `cover_img`) VALUES (1,'Master',1,'2010-02-19','2020-09-13','https://i.ytimg.com/vi/j42DsqccJec/maxresdefault.jpg'),(2,'Main',2,NULL,NULL,'https://images.maariv.co.il/image/upload/f_auto,fl_lossy/c_fill,g_faces:center,h_380,w_500/465598'),(3,'אתמול היה טוב',2,NULL,NULL,'https://i.ytimg.com/vi/sONVbVQXVQM/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&rs=AOn4CLB0A3I3R16fnPCAwDtsQBou9ukuVw'),(5,'El Dorado',3,'2017-05-26','2020-09-14','https://www.mobilemarketingmagazine.com/wp-content/uploads/posts/Shakira_El_Dorado_Campaign.jpg'),(6,'Sale el Sol',3,'2011-02-08','2020-09-17','https://images-na.ssl-images-amazon.com/images/I/81k6yEdpyuL._SL1500_.jpg'),(7,'הפרויקט של עידן רייכל',6,'2002-06-13','2020-09-17','https://images.globes.co.il/Images/NewGlobes/Misc_2/e18_idan-raiher_M.jpg');
/*!40000 ALTER TABLE `albums` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `artists`
--

DROP TABLE IF EXISTS `artists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `artists` (
                           `artist_ID` int NOT NULL AUTO_INCREMENT,
                           `name` varchar(40) NOT NULL,
                           `description` text,
                           `created_at` date DEFAULT NULL,
                           `upload_at` date DEFAULT NULL,
                           `cover_img` varchar(400) DEFAULT NULL,
                           PRIMARY KEY (`artist_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `artists`
--

LOCK TABLES `artists` WRITE;
/*!40000 ALTER TABLE `artists` DISABLE KEYS */;
INSERT INTO `artists` (`artist_ID`, `name`, `description`, `created_at`, `upload_at`, `cover_img`) VALUES (1,'Omer Adam','עומר אדם (נולד ב-22 באוקטובר 1993) הוא זמר ישראלי בסוגת פופ ים-תיכוני.',NULL,NULL,'https://yt3.ggpht.com/a/AATXAJyAFDy5GNROVEvTcg6GapKHn-ws6fVTjzXgxhWg=s176-c-k-c0x00ffffff-no-rj'),(2,'Shlomo Artzi','שלמה ארצי (נולד ב-26 בנובמבר 1949) הוא זמר-יוצר, מוזיקאי, מלחין, פזמונאי, בעל טור ושחקן ישראלי. בין הזמרים המצליחים בישראל ובזמר העברי.',NULL,NULL,'https://yt3.ggpht.com/a/AATXAJwG4HdR6moxBrAgvbFLn2yLGLcQDZzhTFyGhlRo=s144-c-k-c0xffffffff-no-rj-mo'),(3,'Shakira','Shakira Isabel Mebarak Ripoll (born 2 February 1977) is a Colombian singer, songwriter, record producer, dancer, actress, and philanthropist.','1977-02-02','2020-09-13','https://yt3.ggpht.com/a/AATXAJzfa4j62sGChBaAh6R6Je3aCs7qRuswu3Eu9T3XWF4=s144-c-k-c0xffffffff-no-rj-mo'),(6,'עידן רייכל','עידן רייכל (נולד ב-12 בספטמבר 1977) הוא מוזיקאי, פזמונאי, מלחין, זמר-יוצר, מעבד ומפיק מוזיקלי ישראלי. מוכר בעיקר כמפיק \"הפרויקט של עידן רייכל\", שחמשת אלבומיו הראשונים מכרו מעל לחצי מיליון עותקים.[1] רייכל נחשב לאחד האמנים הבולטים והמצליחים בישראל ולאחד האמנים הישראלים המוכרים בעולם.','2011-08-03','2020-09-17','https://yt3.ggpht.com/a/AGF-l7-6fgKaHXdAamKJnnak1zw2-8yRf2CLc-x52g=s800-c-k-c0xffffffff-no-rj-mo');
/*!40000 ALTER TABLE `artists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `interactions`
--

DROP TABLE IF EXISTS `interactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `interactions` (
                                `interaction_ID` bigint unsigned NOT NULL AUTO_INCREMENT,
                                `user_ID` int NOT NULL,
                                `song_ID` int NOT NULL,
                                `liked` tinyint(1) DEFAULT NULL,
                                `play_duration` datetime NOT NULL,
                                `created_at` date DEFAULT NULL,
                                `upload_at` date DEFAULT NULL,
                                UNIQUE KEY `interaction_ID` (`interaction_ID`),
                                KEY `user_ID` (`user_ID`),
                                KEY `song_ID` (`song_ID`),
                                CONSTRAINT `interactions_ibfk_1` FOREIGN KEY (`user_ID`) REFERENCES `users` (`user_ID`),
                                CONSTRAINT `interactions_ibfk_2` FOREIGN KEY (`song_ID`) REFERENCES `songs` (`song_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `interactions`
--

LOCK TABLES `interactions` WRITE;
/*!40000 ALTER TABLE `interactions` DISABLE KEYS */;
/*!40000 ALTER TABLE `interactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `playlists`
--

DROP TABLE IF EXISTS `playlists`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `playlists` (
                             `playlist_ID` int NOT NULL AUTO_INCREMENT,
                             `name` varchar(40) NOT NULL,
                             `created_at` date DEFAULT NULL,
                             `upload_at` date DEFAULT NULL,
                             `cover_img` varchar(400) DEFAULT NULL,
                             `user_ID` int NOT NULL,
                             PRIMARY KEY (`playlist_ID`),
                             KEY `user_ID` (`user_ID`),
                             CONSTRAINT `playlists_ibfk_1` FOREIGN KEY (`user_ID`) REFERENCES `users` (`user_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlists`
--

LOCK TABLES `playlists` WRITE;
/*!40000 ALTER TABLE `playlists` DISABLE KEYS */;
INSERT INTO `playlists` (`playlist_ID`, `name`, `created_at`, `upload_at`, `cover_img`, `user_ID`) VALUES (1,'David Songs','2020-09-10','2020-09-14','https://blog.brightonps.sa.edu.au/wp-content/uploads/2019/02/circle-made-music-instruments_23-2147509304.jpg',2),(2,'runing','2020-09-10','2020-09-14','https://jdpodiatry.com.au/wp-content/uploads/2016/05/woman-runing-sunrise.jpg',2),(4,'Quiet','2020-09-17','2020-09-17','https://i.pinimg.com/originals/0d/9c/df/0d9cdf4892862d0178f8e2db7c845ab9.jpg',2);
/*!40000 ALTER TABLE `playlists` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `playlists_songs`
--

DROP TABLE IF EXISTS `playlists_songs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `playlists_songs` (
                                   `song_ID` int NOT NULL,
                                   `playlist_ID` int NOT NULL,
                                   PRIMARY KEY (`song_ID`,`playlist_ID`),
                                   KEY `playlist_ID` (`playlist_ID`),
                                   CONSTRAINT `playlists_songs_ibfk_1` FOREIGN KEY (`song_ID`) REFERENCES `songs` (`song_ID`),
                                   CONSTRAINT `playlists_songs_ibfk_2` FOREIGN KEY (`playlist_ID`) REFERENCES `playlists` (`playlist_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `playlists_songs`
--

LOCK TABLES `playlists_songs` WRITE;
/*!40000 ALTER TABLE `playlists_songs` DISABLE KEYS */;
INSERT INTO `playlists_songs` (`song_ID`, `playlist_ID`) VALUES (1,1),(2,1),(7,1),(13,1),(1,2),(10,2),(6,4),(7,4),(9,4),(10,4),(11,4);
/*!40000 ALTER TABLE `playlists_songs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `songs`
--

DROP TABLE IF EXISTS `songs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `songs` (
                         `song_ID` int NOT NULL AUTO_INCREMENT,
                         `name` varchar(40) DEFAULT NULL,
                         `album_ID` int NOT NULL,
                         `artist_ID` int NOT NULL,
                         `lyrics` text,
                         `length` time DEFAULT NULL,
                         `track_number` int DEFAULT NULL,
                         `created_at` date DEFAULT NULL,
                         `upload_at` date DEFAULT NULL,
                         `youtube_link` text,
                         PRIMARY KEY (`song_ID`),
                         KEY `artist_ID` (`artist_ID`),
                         KEY `album_ID` (`album_ID`),
                         CONSTRAINT `songs_ibfk_1` FOREIGN KEY (`artist_ID`) REFERENCES `artists` (`artist_ID`),
                         CONSTRAINT `songs_ibfk_2` FOREIGN KEY (`album_ID`) REFERENCES `albums` (`album_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `songs`
--

LOCK TABLES `songs` WRITE;
/*!40000 ALTER TABLE `songs` DISABLE KEYS */;
INSERT INTO `songs` (`song_ID`, `name`, `album_ID`, `artist_ID`, `lyrics`, `length`, `track_number`, `created_at`, `upload_at`, `youtube_link`) VALUES (1,'לעולם בעקבות השמש',1,1,'יום חמים,\nיום קסמים, \nיום תמים\nהלכנו עם השמש\nבערוב,\nיום זהוב,\nלילה טוב\nהיה לנו ליל אמש\n\nבוקר קם מול הים,\nאי משם \nנכנס היום בשער \nבחלון - עץ אלון \nווילון ניצתו באור השחר\n\nכן, כן, כן, כן \nלעולם בעקבות השמש,\nלעולם בעקבות האור.\nהשמש את יומי רושמת\nוליבי ציפור.\n\nבן אדם, \nקום נרדם, \nלא נדם \nהרוח שבתכלת.\nאור קרב, ראש הרם \nאור זורם ישר אל סף הדלת.\n\nכן, כן, כן, כן...','00:03:01',1,'2019-07-04','2019-07-04','https://www.youtube.com/watch?v=Cmmlg6LcYjU&list=RDCmmlg6LcYjU&start_radio=1'),(2,'יש דברים שלא עושים',1,1,'ילא אוכלים לפני ששוטפים ידיים\nלא נכנסים לים המלח אחרי שמגלחים את הרגליים\nלא מקלפים קלמנטינה בחדר שיש בו אנשים\nיש דברים שלא עושים, יש דברים שלא!\nלא עושים מקלחת בלי להשתמש במים\nלא אשכחך ירושלים \nלא מוותרים על אופטומטריסט לפני שקונים משקפיים\nיש דברים שלא עושים, יש דברים שלא!\nחיפשתי משהו מתאים\nבול לפנים\nמשהו יפה\nעיניים תמיד מכבדים\nאז קניתי משקפיים לראות יותר ורוד\nראיתי בעיניים מה גורם לנו לרקוד\nבואו נרים לחיים רואים למרחקים\nכי בחיים האלה יש דברים שלא עושים\nעם אופטומטריסט\nפה במידל איסט\nאתה תראה שש שש\nבסוף אתה תראה את הכביש\nקניתי משקפיים\nזה לא כזה פרויקט\nחבל על העיניים\nאז תיתנו להם ריספקט\n\n\nלא מאיימים על יונה עם יין\nלא נועלים סנדלים על הגרביים\nלא מעלים תמונות של אנשים אוכלים\nיש דברים שלא עושים, יש דברים שלא!\n\n\nחיפשתי משהו מתאים\nבול לפנים\nמשהו יפה\nעיניים תמיד מכבדים\nאז קניתי משקפיים לראות יותר ורוד\nראיתי בעיניים מה גורם לנו לרקוד\nבואו נרים לחיים רואים למרחקים\nכי בחיים האלה יש דברים שלא עושים\nעם אופטומטריסט\nפה במידל איסט\nאתה תראה שש שש\nבסוף אתה תראה את הכביש\nקניתי משקפיים\nזה לא כזה פרויקט\nחבל על העיניים\nאז תיתנו להם ריספקט\n\nניי ניי ניי ניי ניי\nניי ניי ניי ניי נה נה נה נה נה ניי\n\n\nלא מוותרים על אופטומטריסט לפני שקונים משקפיים \nלא מוותרים על אופטומטריסט לפני שקונים משקפיים\n\nיש דברים שלא עושים כשנותנים כבוד לעיניים\nבלי אופטומטריסט לא קונים משקפיים ','00:02:36',2,NULL,NULL,'https://www.youtube.com/watch?v=9zHtxhzM-ys'),(4,'לתת ולקחת',2,2,'ניסיתי ניסיתי מאד,\nלהילחם ברוחות,\nלא סיפרתי לך כמה.\nניתקתי את עצמי מהרוב,\nאך בתוך המציאות,\nחטפתי ג\'ננה.\n\nאדם צריך שתהיה לו מילה,\nקצת מקום בעולם,\nאהבה לא נשכחת.\nוקול אמיתי לתפילה,\nורגע מושלם,\nכדי לתת ולקחת,\nולא...לפחד מהפחד.\n\nניסיתי ניסיתי לצוף,\nיחסי אהבה,\nזה משחק מלוכלך קצת.\nמציפור שוב למדתי לעוף,\nלנחות בשלווה,\nלא ליפול על התחת.\n\nאדם צריך שתהיה לו מילה...\n\nניסיתי כי הייתי צריך,\nלו היית במקומי,\nמה היית אז אמרת?\nתמשיך כי צריך להמשיך,\nאם אתה אמיתי,\nבשבילי אתה המלך.\n\nאדם צריך שתהיה לו מילה...','00:03:48',1,NULL,NULL,'https://www.youtube.com/watch?v=C-bAr0i0YAg'),(6,'ירח',3,2,'וזכרוני הראשון אם יופייך אינו מטעה אותי משליך גיטרה ואבי עליי צורח ואמי אומרת זה החוטא שלי ולוקחת אותי לטייל מנחם מול ירח וזכרוני השני אם עינייך אינן מטעות אותי עוזב את הבית אל סיכוני הזמן ששאלתי אותך האם תהיי לי לעולמים ענית תביט לירח שם יש כבר אדם ואחרי שנה שכרנו חדר מול ירח מתחיל לנגן בבר קודר ברחוב המסגר והם צרחו תורידו את הנמוך עם הטוקסידו ורק אחד אמר תגיד ילד ילד בוא תוריד את הירח בשבילה הייתה תקופה כזאת שהאושר בא בזעם צחקנו מהכל שרפנו את מה שבא ליד לא נשאר לנו אלא לחבק את הצער להגיד אתמול היה טוב ויהיה גם מחר אתמול היה טוב אתמול היה טוב אתמול היה טוב ויהיה גם מחר אתמול היה טוב אתמול היה טוב אתמול היה טוב ויהיה גם מחר קשה לי להתרכז כי יופייך עוד מהמם אותי קשה להגיד חבל או לומר אולי במקום זה אני רוקד וצועק לירח רד ומאשים את העולם בכאביי לפעמים אני שוכח איך התחלתי מול ירח כשאתה שוכח כן כן אתה מסכן יש גשם בשמיים אין ירח בינתיים וכשיצא נלך ליחד עד שנעלם הייתה תקופה כזאת שהאושר בא בזעם צחנו מהכל שרפנו את מה שבא ליד לא נשאר לנו אלא לחבק את הצער להגיד אתמול היה טוב ויהיה גם מחר אתמול היה טוב אתמול היה טוב אתמול היה טוב ויהיה גם מחר אתמול היה טוב אתמול היה טוב אתמול היה טוב ויהיה גם מחר','00:05:42',1,NULL,NULL,'https://www.youtube.com/watch?v=hGiNUsSHkr0'),(7,'ארץ חדשה',2,2,'גשם כבר יורד וזה חורף תל אביב חסומה וגם חיפה. שב ילד שב, אני אומר לך שב ושנינו נוסעים בדרכי עפר. מביט מבעד לזגוגית, יש לנו ארץ למה עוד אחת, בחוץ שקיעה של יום שני, וערבים מתפללים כי איזה חג. חבר למן מסע כזה בחורף מסתכל בי. רגליו קצרות אבל ראשו חכם שנינו במנוסה, הכל פה זז, אומר לי אתה גם אבא, אתה גם בן אדם. מביט מבעד לזגוגית, עיניים יש לו רגישות, כן כן. מוזר איך האויב הזר נראה לו אנושי וגם פוחד. יש לי אישה, זאת אמא שלך ניסע, ניסע, אולי נגיע עד מחר, אם לא נאט, לא נביט, לא נשים לב לפרטים, לא נגיע לארץ חדשה לא נגיע, נגיע, נגיע לארץ חדשה. שתי כבשים עולות על אם הדרך לא נדרוס אותן, אנחנו לא דורסים. שב ילד, שב. אני אומר לך שב. חלב בשפע זה לא אומר ניסים, הוא שולף מצלמה של כיס, הוא חושב שגן העדן מדוייק הוא מכחיש שקר לו שמפחיד ומצלם כדי שנזכור מה שהיה. קראתי בעיתון על אחת בת מאה, שכל חייה עשתה מעשים טובים. שב ילד, שב. אני אומר לך שב. לא כל האנשים נולדו רעים. מזמן, היא מתה די מזמן אביך הוא ימות גם יום אחד. לא, גן העדן לא קיים, אולי קיימת ארץ חדשה. יש לי אישה, זאת אמא שלך... גשם כבר יורד וזה חורף פעם זה הכל היה ורוד, שב ילד, שב למה להתרפק משהו חורק בזיכרון. מביט מבעד לשלטים, עיניים יש כדי להסתכל. תגיד שואל אותי האם, האם יתנו לנו בכלל להיכנס. חברים למן מסע כזה בחורף כבר חסר לי, אבי ישן, זקן ומסתגר. איתו הלכתי דרך העיניים שכבר אין לי עכשיו ילדי שלי איתי הולך. מביט מבעד לזגוגית, יש לנו ארץ - למה עוד אחת? בחוץ שקיעה של יום שני, בפנים אני והוא כמו איש אחד. יש לי אישה, זאת אמא שלך...','00:06:23',2,NULL,NULL,'https://www.youtube.com/watch?v=DRgJz2uRopQ'),(8,'חום יולי אוגוסט',3,2,'חום יולי אוגוסט אז היה כבד מאוד, שעת צהריים, הפלוגה הלכה בוואדי. תרשום בספר דף, מדובר במלחמות, תרשום פצועים שרועדים, וזה נורמלי.  הנעורים יפים, הקיץ אין סופי. ועקנין הגיע לפלוגה, אלוף קריית גת, שרף צריפים, ברח לבית גוברין, המשטרה בעקבותיו, רק פה נרגע קצת.  את מה שאני זוכר מזה אני רושם, מרחפים באלונקות שניים בלי שם, בזמן האחרון אצלי תמונות חוזרות משם, בזיכרון המעומעם שלי היום, מסך כבד של מלחמות והזיות.  חום יולי אוגוסט אז, איצטרובל אחד נושר מסוק נוחת, אני שוכב לבד בשטח. כשהמסוק נעלם, אני פתאום רעב, מת לאכול אותך בבסיס האם לנצח.  אני חוזר עם פס אלייך לביתך, סוגר ת\'חדר, גם אלוהים לא יכנס פה. פתאום אביך נכנס, נראה כמו בוכה, בתעלה פלוגה שלמה חטפה אש נ\"ס ב\"טמפו\".  את מה שאני זוכר מזה אני רושם, תופס מונית חיפנית, קופץ לדיסקוטק. זונות על הגדר אצלי בגוף רק שד בוער, הולך לרקוד עם חיילים מתים בלב. תרשום תרשום, רושם רושם.','00:07:36',2,NULL,NULL,'https://www.youtube.com/watch?v=jlvk5iqhTYw&list=PLiZHIe4ugPFoCCqGo-0X4v6BYUEeizm0Y&index=3'),(9,'כמו אז',3,2,'וכל הדרך, כן, כל הדרך כשהם צחקו, נרדמתי שם. לא מצטער שלא הערת אותי. חלמתי שהכל כמו אז.  עולם מתמוטט, מרוב אמביציות, מרוב שאהבתי, לא אהבת. מסע בדימיון כי שם הכי טוב. חלמתי שהכל כמו אז.  כמו אז מצאתי לי בחושך סיגריה להדליק, כמו אז. כמו אז ישבתי לי בחושך, שם את נמצאת איתי כמו אז.  היה עוד אור, חמש בערך, לא זמן אהבה, לא זמן וידוי, לא, לא שיכור, כרעתי ברך אמרתי: \"לא שמח, לא עצוב, תתחתני איתי שוב?\".  כמו אז נתתי לך בחושך, טבעת ענקית, כמו אז. כמו אז ישבנו שוב בחושך, שם את נמצאת איתי כמו אז.  כמו אז נתתי לך בחושך טבעת ענקית, כמו אז. כמו אז אהבת אותי בחושך, שם את נמצאת איתי כמו אז.  וכל הדרך, כן, כל הדרך, לא מצטער שלא הערת אותי.','00:05:40',3,'2019-07-11',NULL,'https://www.youtube.com/watch?v=tr1D-24HC04&list=PLiZHIe4ugPFoCCqGo-0X4v6BYUEeizm0Y&index=4'),(10,'פעם תורי ופעם תורך',3,2,'לילה אחד כשהלכנו בתוך העיר, עברו בי כמה הרהורים על אנשים שחיים פה, האם הם שונים מאיתנו? הולכים לישון כמו ילדים בעריסות, הולכים עד הסוף. רק אצלנו, חבל, תראי איך אצלנו, פעם תורי ופעם תורך לאהוב.  יכול היה להיות יותר פשוט, שתסתגרי איתי בחדר החשוך, שלא תראי אחר מלבדי, שתהיי שלי.  קניתי לך ג\'ינס חדשים, קנית לי חולצה לבנה שתתאים, פעם תורי לאהוב ופעם תורך. על קיר בעיר כתב מיואש צעיר: \"עיניה פוצעות, אבל נשארתי שלם\".  ובלילה ההוא, כשהלכנו בתוך העיר, דיברתי אלייך ולא ענית לי. מה הפריד אז בינינו? רגש כזה או רגש אחר, או אי הבנה עד הסוף. אצלנו, חבל, תראי איך אצלנו, פעם תורי ופעם תורך לאהוב.  יכול להיות יותר פשוט...  קניתי לך תקליט של האבנים, קנית לי עגיל, אבל זה לא לגילי פעם תורי לאהוב ופעם תורך. על קיר בעיר כתב מיואש צעיר: \"עיניה פוצעות, אבל נשארתי שלם\".  קניתי לך ג\'ינס חדשים...','00:07:10',4,'2019-11-07','2020-09-14','https://www.youtube.com/watch?v=dXzVhwnE6FA&list=PLiZHIe4ugPFoCCqGo-0X4v6BYUEeizm0Y&index=5'),(11,'האהבה הישנה',3,2,'האהבה הישנה, זאת שברחה לי, מונחת שם ואף אחד כבר לא נוגע, מזמן לפני כל הסרטים הפורנוגרפיים, היה עוד משהו פשוט להתגעגע.  הכינורות המתוקים לא מנגנים לי, לא אשקר לך שאצלי הכל בסדר, אין שום פטנטים, העולם נהיה כבר גימיק, אז בואי ונחזור לאינטימיות בחדר.  בשלט רחוק, מביט רחוק, את שוב לובשת לך שמלת כלה, ואני תקוע בחליפת חיינו. בשלט רחוק מביט רחוק, לא אני לא אניח לך, עד שהאהבה, תחזור אלינו.  את יודעת לא הייתי שר לך סתם כך, הנביאים כבר מתו, ואין מי שינבא לי, חלליות עפות למעלה ולמטה, פתאום זה שיר געגועים או שנדמה לי...  בשלט רחוק...','00:03:45',5,'2019-07-11','2020-09-14','https://www.youtube.com/watch?v=of6ClD2-2yU&list=PLiZHIe4ugPFoCCqGo-0X4v6BYUEeizm0Y&index=6'),(12,'Me Enamoré',5,3,'Mi vida me empezó a cambiar La noche que te conocí Tenía poco que perder Y la cosa siguió así Yo con mis sostén a rayas Y mi pelo a medio hacer Pensé: éste todavía es un niño Pero, ¿qué le voy a hacer? Es lo que andaba buscando El doctor recomendando Creí que estaba soñando, oh, oh ¿De qué me andaba quejando? No sé qué estaba pensando Hoy pal cielo voy pateando oh, oh Me enamoré, me ena-ena-namoré Lo vi solito y me lancé Me ena-na-namoré Me ena-na-namo Mira qué cosa bonita Qué boca más redondita Me gusta esa barbita Y bailé hasta que me cansé Hasta que me cansé, bailé Y me ena-na-namoré Nos enamoramos Un mojito, dos mojitos Mira qué ojitos bonitos Me quedo otro ratito Contigo yo tendría diez hijos Empecemos por un par Solamente te lo digo Por si quieres practicar Lo único…','00:03:51',1,'2017-05-12','2020-09-14','https://www.youtube.com/watch?v=sPTn0QEhxds&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo'),(13,'Nada',5,3,'Voy caminando sobre un mar de hojas secas Vuelan los ángeles sobre Berlín Van entonando junto a mí un aleluya Mientras la lluvia cae dentro de mí Extraño tu voz Estoy en tierra de nadie Me falta hasta el aire De espaldas al sol Pasa otro día sin ti No sirve de nada llegar aún más lejos Ni toda la fama, ni todo el dinero No sirve de nada si no estás conmigo Y la soledad se me clava en los huesos No sirve de na, ah, ah, nada Ah, ah, ah, nada Oh oh, nada, no oh Nada No, no, no, no Nadie adivina que depara el destino Nuestro camino aún esta a medio hacer Tiro mi bolso Chanel del barrio chino Al suelo mojado y me lanzo a correr Extraño tu voz Me hace falta tocarte Olerte, mirarte De espaldas al sol Acaba otro día sin ti No sirve de nada llegar aún más lejos Ni toda la…','03:11:00',2,'2018-11-04','2020-09-17','https://www.youtube.com/watch?v=UjX10jO-p3c&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=2'),(14,'Sale el Sol',6,3,'Estas semanas sin verte Me parecieron anos Tanto te quise besar Que me duelen los labios Mira que el miedo nos hizo Cometer estupideces Nos dejo sordos y ciegos Tantas veces Y un día después De la tormenta Cuando menos piensas sale el sol De tanto sumar Pierdes la cuenta Porque uno y uno No siempre son dos Cuando menos piensas Sale el sol Te llore hasta el extremo De lo que era posible Cuando creí lo que era Invencible No hay mal que dure cien anos Ni cuerpo que lo aguante Y lo mejor siempre espera Adelante Y un día después De la tormenta Cuando menos piensas sale el sol De tanto sumar Pierdes la cuenta Porque uno y uno No siempre son dos Cuando menos piensas Sale el sol Cuando menos piensas Sale el sol Y un día después De la…','03:31:00',1,'2011-02-09','2020-09-17','https://www.youtube.com/watch?v=mqqLoUcLX5I&list=OLAK5uy_n78-gMAcT2yb_jY20GyIFPdFd7NAQZcv8&index=1'),(15,'ברכות לשנה החדשה',7,6,'','03:31:00',1,'2019-11-05','2020-09-17','https://www.youtube.com/watch?v=wpy4u3MfNzY&list=OLAK5uy_mFjE9adq53ACHfr9-U7MsAxuAQFIdcnU8'),(16,'בואי',7,6,'בואי תני לי יד ונלך אל תשאלי אותי לאן אל תשאלי אותי על אושר אולי גם הוא יבוא, כשהוא יבוא ירד עלינו כמו גשם  בואי נתחבק ונלך אל תשאלי אותי מתי אל תשאלי אותי על בית אל תבקשי ממני זמן זמן לא מחכה, לא עוצר, לא נשאר...נשאר','00:04:07',2,'2012-09-12','2020-09-18','https://www.youtube.com/watch?v=_uJLsc05ZE4&list=OLAK5uy_mFjE9adq53ACHfr9-U7MsAxuAQFIdcnU8&index=2'),(17,'אם תלך',7,6,'אם תלך מי יחבק אותי ככה  מי ישמע אותי בסוף היום  מי ינחם וירגיע  רק אתה יודע   ואם תלך למי אחכה בחלון בשמלה של חג  שיגיע, יחבק אותי ככה, כמו שאתה מגיע   כשתלך לשמש אצא, בשדה המוזהב, בוקר יום וערב,  ירח יאיר את פני שחולמות כל היום רק עליך   וכשתבוא תישא אותי בשתי ידיך, משדה לנהר,  תרחוץ את פני ותגיד לי מילים  כמו שרק אתה יודע','00:02:50',3,'2003-06-12','2020-09-18','https://www.youtube.com/watch?v=CtpCGfRRToo&list=OLAK5uy_mFjE9adq53ACHfr9-U7MsAxuAQFIdcnU8&index=3'),(18,'אייל אילה',7,6,'','00:03:35',3,'2019-11-05','2020-09-18','https://www.youtube.com/watch?v=xZbj6zRnCg0&list=OLAK5uy_mFjE9adq53ACHfr9-U7MsAxuAQFIdcnU8&index=4'),(19,'קוקוריקו',1,1,'I said I love you baby  You never met a girl like me  Now it\'s time for you to stop  Celebrating your loneliness   Si senyora hey come and play with me  We\'ve got one chance one life  It\'s a matter of faith or destiny  Looking for my luck tonight   Quattro cinque sei sette Say the word and I\'ll be there  Amore Amore Amore חביבי אינשאללה חביבי הלילה חביבי תגידי לי חביבי  תגידי מה קורה TUTTO BENE , תגידי מה קורה לךTUTTO BENE תגידי מה יהיה TUTTO BENE  No se que me esta pasando  I wanna dance dance it for me baby you\'ll see  qué תגיד לי מה מה מה מה  תגידי מה מה מה  Kuku riku Hey senyora (si) come and dance with me  I have all night for you Andiamo Andiamo 1..2..3 Perfecto Quattro cinque sei sette Say the word and I\'ll be there  Amore Amore Amore חביבי אינשאללה חביבי הלילה חביבי תגידי לי חביבי  תגידי מה קורה TUTTO BENE ,תגידי מה קורהTUTTO BENE תגידי מה יהיה TUTTO BENE  No se que me esta pasando I wanna dance dance it for me baby you\'ll see   qué תגיד לי מה מה מה מה  תגידי מה מה מה  Kuku riku','00:02:59',3,'2020-07-02','2020-09-19','https://www.youtube.com/watch?v=d60H5D9GefE'),(20,'תפילה',1,1,'הוא היושב לו אי שם במרומים  הוא הרופא כל חולים  הוא הנותן רוב שמחה לילדים  הוא העושה משפטים  הוא בשמיים והוא היחיד  הוא הגדול הנורא  הוא השומר עלינו מצרה   אלוה, שמור נא עלינו כמו ילדים  שמור נא ואל תעזוב  תן לנו אור ושמחת נעורים  תן לנו כוח עוד ועוד  שמור נא עלינו כמו ילדים  שמור נא ואל תעזוב  תן לנו אור ושמחת נעורים  תן לנו גם לאהוב   מה כבר נותר לנו עוד בימים  מה כבר נותר כל היום  שמש תקווה והמון מבטים  לילה ויום של חלום  הוא בשמיים והוא היחיד  הוא הגדול הנורא  הוא השומר עלינו מצרה   אלוה, שמור נא עלינו כמו ילדים...','00:03:21',4,'2020-04-29','2020-09-19','https://www.youtube.com/watch?v=mQiTfvht20I'),(21,'מתחילים מחדש',1,1,'ותן חלקנו הספר שגרם לאמני ישראל לשתף פעולה ולגלות את האור  ספרון חודשי ללימוד גמרא יומי קצר של 5 דק\' בכל יום  סינגל מתוך פרויקט ״ותן חלקינו״ ״ ואת האור שלך עלי תיתן וגם את כל המחשבות עזור לי לכוון ״  מילים ולחן: תלמידי ותן חלקנו עיבוד והפקה מוזיקלית: מתן דרור  פרוייקט ״ותן חלקינו״ נוצר ממפגש שבועי של אמנים כותבים ומלחינים שיוצאים למסע של שיתוף הקשבה והרבה אמת אחד לשני. הגורם המחבר הוא הספר ״ותן חלקינו״ שיוצא במהדורה חודשית ובו קטעים מהגמרא. כל אמן יביא ליצירה הזו את האמת הקטנה שלו לטקסט ולהגשה וגרם לזה להרגיש שלם. אמני הפרויקט ויוצריו בחרו להקדיש את ההכנסות כולן לפעילות עמותת \"אור בבית\" שמביאה את אור המזוזה לכל דורש ללא עלות.   משתפפים לפי סדר א-ב אייל גולן-איתי לוי-אליעד-בניה ברבי -דוד דאור-דודי בר דוד-הראל סקעת - טליסמן-יואב יצחק-ישי ריבו-ישי לוי-מאור אדרי-מאור תיתון- - משה פרץ-נתן גושן-עדן חסון-עדן מאירי- עידן יניב-פאר טסי-קובי אפללו-קובי פרץ-רותם כהן- רמי קליינשטיין-שלומי שבת. _ מילים:  ועכשיו מתחילים את הכל מחדש אי אפשר להפסיק תן חלקנו רציתי להיות שם ממש חייב להספיק זה קורה לי בבוקר כשחוזרת אלי הנשמה שיצאה מתוכי אוי איזה אושר שהכל כאן עובר לידי  ואת האור שלך עלי תיתן וגם את כל המחשבות עזור לי לכוון וגם אם לפעמים מרגיש בלי קול מודה לך על הכל לבד אנ\'לא יכול את האמת שלך תשלח עלי גם אם הרגשתי שקשה טיפה גדול עלי למד אותי לדעת ולמחול לסלוח על הכל לבד אנ\'לא יכול   ועכשיו מתחילים את הכל מחדש אין לאן למהר תן לי תכח להיות קצת חלש כדי לוותר זה קורה לי בלילה כשכולם ישנים וזה שוב רק אתה ואני אתה שם למעלה ומרגיש אותך כאן על ידי','00:03:37',5,'2020-09-25','2020-09-19','https://www.youtube.com/watch?v=tEkLNLyia2E'),(22,'Sadot Shel Tiruzim',1,1,'מצד אחד לבד הייתי מלך מצד שני איתך אני ישן בצד שני של המיטה השאלות עולות בלי הרף והפרצוף שלך חמוץ מתוק עזבי יוצא לרוץ רחוק קצת לשתוק איך אני יוצא מזה עכשיו  Один 1 Два   2   Три  3 Четыре 4  אם אין אני לי מי לי מתי כבר תעני לי זה לא קרה עם נילי וגם לא עם אחות שלה סיון נסענו ליוון הדלקנו מנגל קצת עשן הוצאתי ת\'גיטרה שרתי שיר ישן נושן אם תזכרי אותי סיון אוי אוי אוי אם תזכרי אותי סיון שדות של תירוצים אני זורע זה לא אני זו את בעיגול פינות עשיתי דוקטורט כי את הכל רוצה לדעת ומי זו התקשרה מאמי יש לך הודעה קטנה מטוני זה אפלטוני איך אני יוצא מזה עכשיו  Один 1 Два   2   Три  3 Четыре 4  אם אין אני לי מי לי מתי כבר תעני לי זה לא קרה עם נילי וגם לא עם אחות שלה סיון נסענו ליוון הדלקנו מנגל קצת עשן הוצאתי ת\'גיטרה שרתי שיר ישן נושן אם תזכרי אותי סיון אוי אוי אוי אם תזכרי אותי סיון','00:03:02',6,'2016-04-16','2020-09-19','https://www.youtube.com/watch?v=YZWd83Z9GlM'),(23,'Chantaje ',5,3,'Cuando estás bien te alejas de mí Te sientes sola y siempre estoy ahí Es una guerra de toma y dame Pues dame de eso que tienes ahí Oye baby no seas mala No me dejes con las ganas Se escucha en la calle Que ya no me quieres Ven y dímelo en la cara Pregúntale a quien tú quieras Vida, te juro que eso no es así Yo nunca tuve una mala intención Yo nunca quise burlarme de ti Conmigo ves, nunca se sabe Un día digo que no y otro que sí Yo soy masoquista Con mi cuerpo, un egoísta Tú eres puro, puro chantaje Puro, puro chantaje Siempre es a tu manera Yo te quiero aunque no quieras Tú eres puro, puro chantaje Puro, puro chantaje Vas libre como el aire No soy de ti ni de nadie Como tú me tientas cuando tú te mueves Ese movimiento sexy siempre me…','00:03:19',3,'2016-11-18','2020-09-19','https://www.youtube.com/watch?v=6Mgqbai3fKo&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=3'),(24,'Miami lauch',5,3,'Lucky you were born that far away so We could both make fun of distance Lucky that I love a foreign land for The lucky fact of your existence Baby, I would climb the Andes solely To count the freckles on your body Never could imagine there were only Ten million ways to love somebody  [Pre-Chorus] Le ro lo le lo le, le ro lo le lo le Can\'t you see? I\'m at your feet  [Chorus] Whenever, wherever We\'re meant to be together I\'ll be there and you\'ll be near And that\'s the deal, my dear Thereover, hereunder You\'ll never have to wonder We can always play by ear But that\'s the deal, my dear  [Verse 2] Lucky that my lips not only mumble They spill kisses like a fountain Lucky that my breasts are small and humble So you don\'t confuse them with mountains Lucky I have strong legs like my mother To run for cover when I need it And these two eyes that for no other The day you leave will cry a river  [Pre-Chorus] Le ro lo le lo le, le ro lo le lo le At your feet I\'m at your feet  [Chorus] Whenever, wherever We\'re meant to be together I\'ll be there and you\'ll be near And that\'s the deal, my dear Thereover, hereunder You\'ll never have to wonder We can always play by ear But that\'s the deal, my dear  [Bridge] Le ro le le lo le, le ro le le lo le Think out loud, say it again Le ro lo le lo le lo le Tell me one more time That you\'ll live Lost in my eyes  [Chorus] Whenever, wherever We\'re meant to be together I\'ll be there and you\'ll be near And that\'s the deal, my dear Thereover, hereunder You\'ve got me head over heels There\'s nothing left to fear If you really feel the way I feel  Whenever, wherever We\'re meant to be together I\'ll be there and you\'ll be near And that\'s the deal, my dear Thereover, hereunder You\'ve got me head over heels There\'s nothing left to fear If you really feel the way I feel','00:02:04',4,'2017-05-30','2020-09-19','https://www.youtube.com/watch?v=AgLLpBCNI6E&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=4'),(25,'Amarillo ',5,3,'¿Qué será lo que fabricas tú? ¿Lo que te hace tan especial? Cuando hablo o pienso en tí Es imposible ser imparcial No es un daño de la percepción Tampoco falta de sentido común Te haz vuelto todo y más alla El mundo lo sabe, no es ningún tabú Es tan irreal Lo que me hace extrañarte así cuando no estás Lo que me hace idolatrarte así, es todo tú Estoy perdida, vivo anclada casi como un pájaro a su cielo azul Te estimo Amarillo, me tienes en los bolsillos Morado, ya me olvidé del pasado En rojo, por que me sangran los ojos de llorarte Cuando no estás a mi lado Celeste, cuésteme lo que me cueste Dorado, por que no pienso perderte Tu amor es un arco iris de colores Y me muero por tenerte Es delicioso verte llegar Ver los efectos que causas, tú Con tu sonrisa angelical…','00:03:41',5,'2017-05-26','2020-09-19','https://www.youtube.com/watch?v=QoNIR8bGSEg&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=5'),(26,'Perro Fiel',5,3,'Aquí estás Ya no puedes detenerte ¿Dónde vas? Si estoy loco por tenerte Cómo lo iba a saber Que te vería otra vez Tú me confundes, no sé qué hacer Yo lo que quiero es pasarla bien Yo tengo miedo de que me guste Y que vaya a enloquecer Si eso pasa yo seguiré Contigo aquí como un perro fiel Yo tengo miedo de que me guste Y que vaya a enloquecer Te hablo en serio mai, no estoy jugando Tanto tiempo pasa y nada Estas ganas no me aguanto Y aunque tú me esquives, yo te sigo deseando Dicen que tú eres peligrosa No le hago caso a esas cosas Dime qué está pasando Me tienes como un loco, soy un loco enamorado, eh Quiero saber cuánto me vas a insistir Y hasta dónde llegarías por mí Siento mucho la espera Pero valdrá la pena cuando te esté besando De la manera que te mueves así…','00:03:15',6,'2017-09-15','2020-09-19','https://www.youtube.com/watch?v=SHq2qrFUlGY&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=6'),(27,'Trap ',5,3,'Quiere que lo hagamos en diferentes partes Pero estoy cansada de desilusiones Hace mucho tiempo no creo en los hombres Y no necesito de este mal de amores Me pide y yo le doy Sabe que siempre aquí estoy Casi siempre llama tarde y nunca cambia el dirty boy Quiero besarte Satisfacerte Oye baby no me niegues, vámonos Quiere que se lo haga en diferentes partes Ella está cansada de desilusiones No quiere saber de un rompecorazones Llámame cuando quieras beba Quiere que lo hagamos en diferentes partes Pero estoy cansada de desilusiones No quiero saber de un rompecorazones Llámame cuando quieras beba Tu muévete encima de mí Compláceme Ay bebé Tu muévete encima de mí Compláceme Cuando tu quieras baby Vámonos, perdámonos De la realidad escapémonos En la cama tu y yo…','00:03:21',7,'2018-01-26','2020-09-19','https://www.youtube.com/watch?v=zkG4Xpz6t68&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=7'),(28,'Comme moi',5,3,'Quiere que lo hagamos en diferentes partes Pero estoy cansada de desilusiones Hace mucho tiempo no creo en los hombres Y no necesito de este mal de amores Me pide y yo le doy Sabe que siempre aquí estoy Casi siempre llama tarde y nunca cambia el dirty boy Quiero besarte Satisfacerte Oye baby no me niegues, vámonos Quiere que se lo haga en diferentes partes Ella está cansada de desilusiones No quiere saber de un rompecorazones Llámame cuando quieras beba Quiere que lo hagamos en diferentes partes Pero estoy cansada de desilusiones No quiero saber de un rompecorazones Llámame cuando quieras beba Tu muévete encima de mí Compláceme Ay bebé Tu muévete encima de mí Compláceme Cuando tu quieras baby Vámonos, perdámonos De la realidad escapémonos En la cama tu y yo…','00:03:08',8,'2017-05-31','2020-09-19','https://www.youtube.com/watch?v=YVaahuLdrng&list=OLAK5uy_mLAIIooZEcieJ0n0XKRm6fu-i0khGktBo&index=8'),(29,'Loca',6,3,'Loca (loca) No te pongas bruto, (loca) Que te la bebe Dance or die (loca) El está por mi y por ti borró Y eso que tu tienes to\' Y yo ni un kiki El está por mi Y por ti borro (borro) Y eso que tu tienes to\' Y yo ni un kiki Ella se hace la bruta pa cotizar Asi conmigo enfrente ella se hace la gata en celo contigo Te cotorrea el oído pa tenerte en alta Ella muere por ti, tu por mi es que matas. Sigo tranquila como una paloma de equina Mientras ella se pasa en su BM Mira yo de aquí no me voy, se que esta esta por mi Y ninguna va poder quitármelo de un tirón Yo soy loca con mi tigre Loca, loca, loca Soy loca con mi tigre Loca, loca, loca Soy loca con mi tigre (Loca, loca, loca) Soy loca con mi tigre (Loca, loca, loca) El esta por mi Y por ti borro (borro) Eso que tu tienes…','00:03:03',2,'2017-01-25','2020-09-19','https://www.youtube.com/watch?v=p_uBMc7mA0c&list=OLAK5uy_n78-gMAcT2yb_jY20GyIFPdFd7NAQZcv8&index=2'),(30,'Antes de las Seis',6,3,'No actúes tan extraño Duro como una roca Si te mostré pedazos de piel Que la luz del sol aún no toca Y tantos lunares que ni yo misma conocía Te mostré mi fuerza bruta Mi talón de Alquiles, mi poesía Que harás solo una historia más Que haré si no te vuelvo a ver Oh, oh Si desde el día en que no estás Vi la noche llegar mucho antes de las seis Si desde el día en que no estás Vi la noche llegar mucho antes de las seis Mucho antes No dejes el barco Tanto antes de que zarpemos Hacia alguna isla desierta Y después, después veremos Si me ves desarmada Porque lanzas tus misiles Si ya conoces mis puntos cardinales Los más sensibles y sútiles Que…','00:02:54',3,'2017-01-26','2020-09-19','https://www.youtube.com/watch?v=D4KhLZ9SAzU&list=OLAK5uy_n78-gMAcT2yb_jY20GyIFPdFd7NAQZcv8&index=3'),(31,'Gordita',6,3,'Esto a ti te va a calentar ¿Qué? A calentar ¿Qué? A calentar Esto a ti, a ti te va a calentar ¿Qué? A calentar ¿Qué? A calentarEso que me dices tú, eso Son palabras bonitas, dedicadas Llenas de amor para mí (Tú eres mi gordita) Eso que me dices tú, eso Son palabras bonitas, dedicadas Llenas de amor para míOye, oye Perversión absoluta en el área Perverso como tener sexo en una funeraria Yo soy el jefe; y yo la secretaria Deja que te saque tu lado de ordinaria Shaki, tú estás bien bonita Aunque también me gustabas cuando estabas más gordita Con el pelito negrito y la cara redondita Así, medio rockeritaTambién me gustas ahora pero cuando pierdes los modales Cuando comes sin cubiertos como los animales Cuando se te sale lo sucio, lo obsceno Todo lo que sabe bueno, todo el venenoSi yo soy criminal, tú eres una delincuente Vamos a tirarnos los dos del mismo puente Sin paracaídas en nuestro mundo, volando Como los hippies cuando están fumandoEso que me dices tú, eso Son palabras bonitas, dedicadas Llenas de amor para mí (Tú eres mi gordita)Esto a ti te va a calentar ¿Qué? A calentar ¿Qué? A calentar Esto a ti, a ti te va a calentar ¿Qué? A calentar ¿Qué? A calentarMe gusta que me mires como un delincuente Soy Caperucita, y tú eres el lobo que miente Invítame a pasearme por el lado salvaje Donde yo no tenga que llevar maquillajeNo puedo decirte todo, todo lo que pienso Soy peor que tú, y si me buscas, yo te encuentroChiquita rica que me rasca cuando pica, chica Chiquito, dame un pico chico en el hocico rico Yo poco a poco (poco) Te toco un poco (poco) Tárara, tárara, tara-táraraChiquita rica que me rasca cuando pica, chica Chiquito, dame un pico chico en el hocico rico Yo poco a poco (poco) Te toco un poco (poco) Tárara, tárara, tara-táraraEsto a ti te va a calentar ¿Qué? A calentar ¿Qué? A calentar Esto a ti, a ti te va a calentar ¿Qué? A calentar ¿Qué? A calentarEso que me dices tú, eso Son palabras bonitas, dedicadas Llenas de amor para mí (Tú eres mi gordita) Eso que me dices tú, eso Son palabras bonitas, dedicadas Llenas de amor para mí','00:03:25',3,'2017-01-26','2020-09-19','https://www.youtube.com/watch?v=jIoCKvW2JTU&list=OLAK5uy_n78-gMAcT2yb_jY20GyIFPdFd7NAQZcv8&index=4'),(32,' Addicted to You',6,3,'Debe ser el perfume que usas o el agua con la que te bañas Pero cada cosita que haces a mí me parece una hazaña Me besaste esa noche cual si fuera el único día de tu boca Cada vez que me acuerdo yo siento en mi pecho el peso de una roca Son tus ojos marrones Esa veta verdosa Es tu cara de niño Y esa risa nerviosa I\'m addicted to you Porque es un vicio tu piel Baby I\'m addicted to you Quiero que te dejes querer I\'m addicted to you Porque es un vicio tu piel Baby I\'m addicted to you Quiero que te dejes querer Por el puro placer de flotar ahora sí me llevó la corriente Ya no puedo dormir ni comer como lo hace la gente decente Tu recuerdo ha quedado así como un broche prendido a mi almohada Y tú en cambio que tienes memoria…','00:02:34',4,'2012-05-02','2020-09-19','https://www.youtube.com/watch?v=MntbN1DdEP0&list=OLAK5uy_n78-gMAcT2yb_jY20GyIFPdFd7NAQZcv8&index=5'),(33,'Lo Que Más',6,3,'Cuántas veces nos salvó el pudor Y mis ganas de siempre buscarte Pedacito de amor delirante Colgado de tu cuello un sábado de lluvia a la cinco de la tarde Sabe Dios, cómo me cuesta dejarte Y te miro mientras duermes, más no voy a despertarte Es que hoy se me agotó la esperanza Porque con lo que nos queda de nosotros ya no alcanza Eres lo que más he querido En la vida lo que mas he querido Eres lo que más he querido en la vida Lo que más he querido Cuántas veces quise hacerlo bien Y pequé por hablar demasiado No saber dónde, cómo, ni cuándo Todos estos años caminando juntos Ahora no parecen tanto Sabe Dios, todo el amor…','00:02:26',5,'2017-01-26','2020-09-19','https://www.youtube.com/watch?v=ZTMbpMKe7fA&list=OLAK5uy_n78-gMAcT2yb_jY20GyIFPdFd7NAQZcv8&index=6'),(34,'מדברים בשקט',7,6,'דבר אלי מילים פשוטות כמו שהיינו פעם כשעוד כתבת מכתבים של אהבה כמו שהיית מחבק נשאר איתי ולא בורח היית בא מביא לי פרח  דבר אלי מילים פשוטות כמו שהיינו פעם כשעוד היית מסתכל עלי נרדמת לצידך כמו שהיית דואג אם תמיד אהיה שלך  פעם היינו מדברים לאט היינו מדברים בשקט מדברים מעט היית מסתכל לי בעיניים כשדיברתי איתך  דבר אלי מילים פשוטות כמו שהיינו פעם כשעוד היית מסתכל עלי נרדמת לצידך כמו שהיית דואג אם תמיד אהיה שלך  פעם היינו מדברים לאט...  תחבק אותי חזק ואל תרפה ימים קשים עוברים אני נופלת מהכובד על כתפי אני נופלת מהכובד על כתפי  פעם היינו מדברים לאט...','00:03:55',4,'2002-12-11','2020-09-19','https://www.youtube.com/watch?v=J_g6c6IZ6VY&list=PL-XGrtXcEC4HYlMy85pfMrsfAga2gpW40&index=2');
/*!40000 ALTER TABLE `songs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
                         `user_ID` int NOT NULL AUTO_INCREMENT,
                         `name` varchar(40) NOT NULL,
                         `email` varchar(100) NOT NULL,
                         `password` varchar(40) NOT NULL,
                         `is_admin` tinyint(1) NOT NULL DEFAULT '0',
                         `preferences` text,
                         `remember_token` tinyint(1) DEFAULT '0',
                         `created_at` date DEFAULT NULL,
                         `upload_at` date DEFAULT NULL,
                         `role` longtext,
                         `group_type` int DEFAULT NULL,
                         `status` longtext,
                         PRIMARY KEY (`user_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` (`user_ID`, `name`, `email`, `password`, `is_admin`, `preferences`, `remember_token`, `created_at`, `upload_at`, `role`, `group_type`, `status`) VALUES (2,'david','david@gmail.com','david12345',1,'like',0,NULL,NULL,NULL,NULL,NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-11-21  1:55:47
