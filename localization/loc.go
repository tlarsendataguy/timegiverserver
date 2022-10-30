package localization

type Lang int

const (
	DE Lang = iota
	EN
	ES
	FR
	HI
	JA
	PT
	RU
	ZH
)

type StepText struct {
	Title       string
	Description string
}

var NoCaffeine = []StepText{
	{`Kein Koffein`, `Unterlasse Koffeinkonsum. Zusätzlich zu Kaffee und Tee kann Koffein auch in Erfrischungsgetränken und Schokolade vorhanden sein.`},
	{`No caffeine`, `Refrain from consuming caffeine.  In addition to coffee and tea, caffeine may also be present in soft drinks and chocolate.`},
	{`No tomes cafeína`, `Evita consumir cafeína. Además del café y el té, los refrescos y el chocolate también pueden contener cafeína.`},
	{`Pas de caféine`, `Abstenez-vous de consommer de la caféine. En plus du café et du thé, la caféine peut aussi se trouver dans les boissons gazeuses et le chocolat.`},
	{`कोई कैफीन नहीं`, `कैफीन खाने से बचें। कॉफी और चाय के अलावा, कैफीन सॉफ्ट ड्रिंक और चॉकलेट में भी हो सकते हैं।`},
	{`カフェインなし`, `カフェインを摂取しないようにしましょう。コーヒーやお茶の他にも、カフェインはソフトドリンクやチョコレートに含まれている場合があります。`},
	{`Sem Cafeína`, `Evite consumir cafeína. Além de café e chá, a cafeína também pode estar presente em refrigerantes e chocolate.`},
	{`Никакого кофеина`, `Воздержитесь от кофеина. Кроме кофе и чая он содержится в газировке и шоколаде.`},
	{`无咖啡因`, `尽量不摄入咖啡因。除了咖啡和茶，软饮料和巧克力也可能含有咖啡因。`},
}

var CaffeineOk = []StepText{
	{`Koffein Ok`, `Koffeinkonsum ist nach deinem Ermessen erlaubt.`},
	{`Caffeine Ok`, `Caffeine consumption is allowed at your discretion.`},
	{`Puedes tomar cafeína`, `El consumo de cafeína está permitido en la cantidad que consideres adecuada.`},
	{`Caféine autorisée`, `Vous pouvez consommer autant de caféine que vous le souhaitez.`},
	{`कैफीन ठीक`, `कैफीन खपत की अनुमति आपके विवेकाधीन है।`},
	{`カフェインOK`, `カフェインは自己の責任で摂取してください。`},
	{`Cafeína Ok`, `O consumo de cafeína é permitido a seu critério.`},
	{`Можно кофеин`, `Кофеин можете употреблять по своему усмотрению.`},
	{`咖啡因正常`, `您可以自行决定是否摄入咖啡因。`},
}

var Caffeine2C = []StepText{
	{`Trinke ein koffeinhaltiges Getränk`, `Trinke 1 - 2 Tassen schwarzen Kaffee, Tee oder Entsprechendes.`},
	{`Drink a caffeinated beverage`, `Drink 1-2 cups of black coffee, tea, or equivalent.`},
	{`Toma una bebida con cafeína`, `Bebe 1 o 2 tazas de café solo, té o algo similar.`},
	{`Buvez une boisson caféinée`, `Buvez 1 à 2 tasses de café noir, thé, ou équivalent.`},
	{`कोई कैफीन वाला पेय पिएं`, `1-2 कप ब्लैक कॉफ, चाय या उसी मात्रा में कुछ और पिएं।`},
	{`カフェイン入り飲料を飲む`, `1～2杯のブラックコーヒー、お茶、または類似した飲料を飲みましょう。`},
	{`Beba uma bebida cafeinada`, `Beba 1-2 xícaras de café preto, chá, ou equivalente.`},
	{`Выпейте напиток с кофеином`, `Выпейте 1-2 чашки черного кофе, чая, или их эквивалент.`},
	{`喝咖啡因饮料`, `喝1-2杯黑咖啡、茶或等价物。`},
}

var Caffeine3C = []StepText{
	{`Trinke ein koffeinhaltiges Getränk`, `Trinke 2 - 3 Tassen schwarzen Kaffee, Tee oder Entsprechendes.`},
	{`Drink a caffeinated beverage`, `Drink 2-3 cups of black coffee, tea, or equivalent.`},
	{`Toma una bebida con cafeína`, `Bebe 2 o 3 tazas de café solo, té o algo similar.`},
	{`Buvez une boisson caféinée`, `Buvez 1 à 2 tasses de café noir, thé, ou équivalent.`},
	{`कोई कैफीन वाला पेय पिएं`, `2-3 कप ब्लैक कॉफ, चाय या उसी मात्रा में कुछ और पिएं।`},
	{`カフェイン入り飲料を飲む`, `2～3杯のブラックコーヒー、お茶、または類似した飲料を飲みましょう。`},
	{`Beba uma bebida cafeinada`, `Beba 2-3 xícaras de café preto, chá, ou equivalente.`},
	{`Выпейте напиток с кофеином`, `Выпейте 2-3 чашки черного кофе, чая, или их эквивалент.`},
	{`喝咖啡因饮料`, `喝2-3杯黑咖啡、茶或等价物。`},
}

var LightBreakfast = []StepText{
	{`Esse ein leichtes Frühstück`, `Iss ein leichtes, proteinreiches Frühstück. Die Mahlzeit sollte idealerweise zwischen 200 und 300 Kalorien enthalten.`},
	{`Eat a light breakfast`, `Eat a light, high-protein breakfast.  Ideally, the meal should contain between 200 and 300 calories.`},
	{`Come un desayuno ligero`, `Come un desayuno ligero y alto en proteínas. Lo ideal es que contenga entre 200 y 300 calorías.`},
	{`Mangez un petit déjeuner léger`, `Mangez un petit déjeuner léger et riche en protéines. Idéalement, le repas devrait contenir entre 200 et 300 calories.`},
	{`कोई हल्का नाश्ता लें`, `हल्का, उच्च प्रोटीन युक्त नाश्ता करें। आदर्श रूप से, खाना 200 से 300 कैलोरी के बीच होना चाहिए।`},
	{`軽い朝食を食べる`, `タンパク質が多く含まれた軽い朝食（200～300カロリー）を食べましょう。`},
	{`Coma um café da manhã leve`, `Coma um café da manhã leve, rico em proteínas. Idealmente, a refeição deve conter entre 200 e 300 calorias.`},
	{`Съешьте легкий завтрак`, `Съешьте легкий богатый белком завтрак. В идеале в нём должно быть 200-300 калорий.`},
	{`吃简便早餐`, `吃一顿简便、富含蛋白质的早餐。理想情况下，这顿饭的热量应介于200至300卡路里之间。`},
}

var LightLunch = []StepText{
	{`Esse ein leichtes Mittagessen`, `Iss ein leichtes, proteinreiches Mittagessen. Die Mahlzeit sollte idealerweise zwischen 200 und 300 Kalorien enthalten.`},
	{`Eat a light lunch`, `Eat a light, high-protein lunch.  Ideally, the meal should contain between 200 and 300 calories.`},
	{`Come un almuerzo ligero`, `Come un almuerzo ligero y alto en proteínas. Lo ideal es que contenga entre 200 y 300 calorías.`},
	{`Mangez un déjeuner léger`, `Mangez un déjeuner léger et riche en protéines. Idéalement, le repas devrait contenir entre 200 et 300 calories.`},
	{`िन का भोजन हल्का करें`, `दिन का भोजना हल्का उच्च प्रोटीन युक्त होना चाहिए। आदर्श रूप से, खाना 200 से 300 कैलोरी के बीच होना चाहिए।`},
	{`軽い昼食を食べる`, `タンパク質が多く含まれた軽い昼食（200～300カロリー）を食べましょう。`},
	{`Coma um almoço leve`, `Coma um almoço leve, rico em proteínas. Idealmente, a refeição deve conter entre 200 e 300 calorias.`},
	{`Съешьте легкий обед`, `Съешьте легкий богатый белком обед. В идеале в нём должно быть 200-300 калорий.`},
	{`吃简便午餐`, `吃一顿简便、富含蛋白质的午餐。理想情况下，这顿饭的热量应介于200至300卡路里之间。`},
}

var LightDinner = []StepText{
	{`Esse ein leichtes Abendessen`, `Iss ein leichtes Abendessen mit vielen Kohlenhydraten. Die Mahlzeit sollte idealerweise zwischen 300 und 400 Kalorien enthalten.`},
	{`Eat a light dinner`, `Eat a light, high-carbohydrate dinner.  Ideally, the meal should contains between 300 and 400 calories.`},
	{`Come una cena ligera`, `Come una cena ligera y alta en carbohidratos. Lo ideal es que contenga entre 300 y 400 calorías.`},
	{`Mangez un dîner léger`, `Mangez un dîner léger et riche en protéines. Idéalement, le repas devrait contenir entre 300 et 400 calories.`},
	{`रात भोजन हल्का करें`, `रात का खाना हल्का उच्च कार्बोहाइड्रेट युक्त होना चाहिए। आदर्श रूप से, खाना 300 से 400 कैलोरी के बीच होना चाहिए।`},
	{`軽い夕食を食べる`, `炭水化物が多く含まれた軽い夕食（300～400カロリー）を食べましょう。`},
	{`Coma um jantar leve`, `Coma um jantar leve, rico em carboidratos. Idealmente, a refeição deve conter entre 300 e 400 calorias.`},
	{`Съешьте легкий ужин`, `Съешьте легкий богатый углеводами ужин. В идеале в нём должно быть 300-400 калорий.`},
	{`吃简便晚餐`, `吃一顿简便、富含碳水化合物的晚餐。理想情况下，这顿饭的热量应介于300至400卡路里之间。`},
}

var LightDinnerOptional = []StepText{
	{`Esse ein leichtes Abendessen`, `Iss ein leichtes Abendessen mit vielen Kohlenhydraten. Die Mahlzeit ist optional, sollte jedoch beim Verzehr idealerweise zwischen 300 und 400 Kalorien enthalten.`},
	{`Eat a light dinner`, `Eat a light, high-carbohydrate dinner.  This meal is optional, but should contain between 300 and 400 calories if consumed.`},
	{`Come una cena ligera`, `Come una cena ligera y alta en carbohidratos. Esta comida es opcional y debería contener entre 300 y 400 calorías.`},
	{`Mangez un dîner léger`, `Mangez un dîner léger et riche en protéines. Ce repas est optionnel, mais devrait contenir entre 300 et 400 calories.`},
	{`रात भोजन हल्का करें`, `रात का खाना हल्का उच्च कार्बोहाइड्रेट वाला होना चाहिए। यह खाना वैकल्पिक है, परंतु यदि खाना हो तो इसे 300 से 400 कैलोरी के बीच होना चाहिए।`},
	{`軽い夕食を食べる`, `炭水化物が多く含まれた軽い夕食を食べましょう。この食事は任意ですが、食べる場合は摂取カロリーが300～400カロリーになるようにしてください。`},
	{`Coma um jantar leve`, `Coma um jantar leve, rico em carboidratos. Idealmente, a refeição deve conter entre 300 e 400 calorias.`},
	{`Съешьте легкий ужин`, `Съешьте легкий богатый углеводами ужин. Это не обязательно, но если хотите поесть, в ужине должно быть 300-400 калорий.`},
	{`吃简便晚餐`, `吃一顿简便、富含碳水化合物的晚餐。这顿饭可吃可不吃，但如果吃，其热量应介于300至400卡路里之间。`},
}

var HeavyBreakfast = []StepText{
	{`Esse ein großes Frühstück`, `Iss ein herzhaftes, proteinreiches Frühstück.`},
	{`Eat a large breakfast`, `Eat a hearty, high-protein breakfast.`},
	{`Come un desayuno abundante`, `Come un desayuno abundante y alto en proteínas.`},
	{`Mangez un petit déjeuner consistant`, `Mangez un petit déjeuner consistant et riche en protéines.`},
	{`सुबह का नाश्ता बड़ा करें`, `अच्छा, उच्च प्रोटीन युक्त नाश्ता करें।`},
	{`朝食をたくさん食べる`, `タンパク質が多く含まれた朝食をたっぷり食べましょう。`},
	{`Coma um café da manhã grande`, `Coma um café da manhã saudável, rico em proteínas.`},
	{`Съешьте плотный завтрак`, `Съешьте плотный богатый белком завтрак.`},
	{`吃丰盛早餐`, `吃一顿丰盛、富含蛋白质的早餐。`},
}

var HeavyLunch = []StepText{
	{`Esse ein großes Mittagessen`, `Iss ein herzhaftes, proteinreiches Mittagessen.`},
	{`Eat a large lunch`, `Eat a hearty, high-protein lunch.`},
	{`Come un almuerzo abundante`, `Come un almuerzo abundante y alto en proteínas.`},
	{`Mangez un déjeuner consistant`, `Mangez un déjeuner consistant et riche en protéines.`},
	{`दिन का खाना बड़ा करें`, `अच्छा, उच्च प्रोटीन युक्त लंच करें।`},
	{`昼食をたくさん食べる`, `タンパク質が多く含まれた昼食をたっぷり食べましょう。`},
	{`Coma um almoço grande`, `Coma um almoço saudável, rico em proteínas.`},
	{`Съешьте плотный обед`, `Съешьте плотный богатый белком обед.`},
	{`吃丰盛午餐`, `吃一顿丰盛、富含蛋白质的午餐。`},
}

var HeavyDinner = []StepText{
	{`Esse ein großes Abendessen`, `Iss ein herzhaftes Abendessen mit vielen Kohlenhydraten.`},
	{`Eat a large dinner`, `Eat a hearty, high-carbohydrate dinner.`},
	{`Come una cena abundante`, `Come una cena abundante y alta en carbohidratos.`},
	{`Mangez un dîner consistant`, `Mangez un dîner consistant et riche en protéines.`},
	{`रात का खाना बड़ा करें`, `रात के खाने में अच्छा, उच्च कार्बोहाइड्रेट युक्त भोजन लें।`},
	{`夕食をたくさん食べる`, `炭水化物が多く含まれた夕食をたっぷり食べましょう。`},
	{`Coma um jantar grande`, `Coma um jantar saudável, rico em carboidratos.`},
	{`Съешьте плотный ужин`, `Съешьте плотный богатый углеводами ужин.`},
	{`吃丰盛晚餐`, `吃一顿丰盛、富含碳水化合物的晚餐。`},
}

var NoSnack = []StepText{
	{`Kein Essen`, `Iss keine Mahlzeiten oder Snacks.`},
	{`No eating`, `Do not eat any meals or snacks.`},
	{`No comas nada`, `No comas ninguna comida o aperitivo.`},
	{`Pas de repas`, `Ne mangez aucun repas ni snacks.`},
	{`कोई खाना नहीं`, `कोई खाना या नाश्ता ना लें॥`},
	{`食事なし`, `食事や間食を食べないでください。`},
	{`Sem comer`, `Não coma nenhuma refeição ou lanches.`},
	{`Никакой еды`, `Не ешьте и не перекусывайте.`},
	{`不进食`, `不要吃饭或吃零食。`},
}

var SetWatch = []StepText{
	{`Uhr stellen`, `Stelle deine Uhr auf die Zielzeit. Du nimmst jetzt den Wechsel vor, um deinen Tag um die Zeitzone des Ziels herum zu organisieren.`},
	{`Set watch`, `Set your watch to destination time.  You are now making the switch to organize your day around the destination timezone.`},
	{`Ajustar el reloj`, `Pon la hora de destino en tu reloj. Al hacer este cambio estarás organizando el día según la zona horaria de tu destino.`},
	{`Régler l'horloge`, `Réglez votre montre sur l'heure de destination. À partir de maintenant, vous allez organiser votre journée en fonction du fuseau horaire de destination.`},
	{`घड़ी सेट करें`, `अपने घड़ी को अपने मंजिल के अनुसार सेट करें। अब आप अपने दिन की शुरूआत अपने मंजिल के टाइमजोन के अनुसार बदल रहे हैं।`},
	{`時計を設定`, `お手持ちの時計を到着地の時間に合わせて設定しましょう。到着地の時間帯に切り替えるための準備を始めます。`},
	{`Ajustar relógio`, `Ajuste seu relógio para o horário de destino. Agora você está fazendo a mudança para organizar o seu dia em torno do fuso horário de destino.`},
	{`Переведите часы`, `Переведите часы на местное время города прилета. Вам нужно подстроить свой день под местное время.`},
	{`设置手表`, `将手表设为目的地时间。您现在正在向按照目的地时区安排日常生活过渡。`},
}

var Sleep = []StepText{
	{`Schlafen`, `Gehe so bald wie möglich schlafen. Benutze falls nötig eine Augenmaske. Wenn du nicht einschlafen kannst, ist es wichtig, sich auszuruhen und von der Arbeit und sozialen Aktivitäten zu befreien.`},
	{`Sleep`, `Go to sleep as soon as possible, using an eye mask if necessary.  If you are unable to fall asleep, then it is important to rest and disengage from work and social activity.`},
	{`Duerme`, `Duérmete lo antes posible y usa un antifaz si lo necesitas. Si no puedes dormirte, es importante que descanses y desconectes del trabajo y de la actividad social.`},
	{`Dormir`, `Allez vous coucher dès que possible, avec un masque pour les yeux si nécessaire. Si vous n\'arrivez pas à vous endormir, alors il est important de vous reposer et de vous éloigner du travail et d\'une activité sociale.`},
	{`सोएं`, `जितना जल्दी हो सके सो जाएं। जरूरत पड़े तो आंख ढंकने का प्रबंध करें। यदि आपको सोने में दिक्कत आ रही है तो इस बात का खयाल रखें कि आपक किसी काम या सामाजिक कार्य में ना लगे हों।`},
	{`睡眠`, `必要な方はアイマスクを使用し、できるだけ速やかに睡眠をとってください。眠れない場合は、仕事や社会活動はやめて休息をとることが重要です。`},
	{`Dormir`, `Vá para a cama o mais rápido possível, usando uma máscara para os olhos, se necessário. Se você não conseguir dormir, então é importante descansar e se desligar do trabalho e atividade social.`},
	{`Поспите`, `Ложитесь спать как можно раньше, при необходимости с маской для сна. Если заснуть не получается, постарайтесь всё же отдохнуть и воздержаться от работы и социальной активности.`},
	{`睡觉`, `尽早睡觉，必要时请使用眼罩。如果如法入睡，那么结束工作和社交活动并进行休息就很重要。`},
}

var NoNap = []StepText{
	{`Kein Schlaf`, `Unterlasse ein Nickerchen, da dadurch später Schlafmuster gestört werden können.`},
	{`No sleeping`, `Refrain from napping as this may disrupt later sleep patterns.`},
	{`No duermas`, `Evita dormir la siesta, ya que puede que eso altere tus patrones de sueño.`},
	{`Ne pas dormir`, `Évitez de faire une sieste car cela risquerait de perturber ultérieurement le rythme de votre sommeil.`},
	{`सोना बिल्कुल नहीं`, `जितना जल्दी हो सके, सो जाएं। जरूरत पड़े तो आंख ढंकने का प्रबंध करें। यदि आपको नींद नहीं आ रही है तो यह महत्वपूर्ण है कि आप स्वयं को काम और सामाजिक सरोकारों से दूर कर लें।`},
	{`睡眠なし`, `後の睡眠パターンを妨害する恐れがあるので、昼寝は控えてください。`},
	{`Sem Dormir`, `Evite cochilar, pois isso pode perturbar os padrões de sono mais tarde.`},
	{`Никакого сна`, `Постарайтесь не дремать, это может нарушить ритмы сна.`},
	{`不睡觉`, `尽量不要小睡，因为这可能会扰乱以后的睡眠规律。`},
}

var Arrive = []StepText{
	{`Ankunftszeit`, `Wenn deine Reisevorbereitungen pünktlich ablaufen, solltest du bald an deinem Ziel ankommen.`},
	{`Arrival time`, `If your travel arrangements are running on time, you should be arriving at your destination.`},
	{`Hora de llegada`, `Si tu viaje está yendo según lo planeado, deberías estar llegando a tu destino.`},
	{`Heure d'arrivée`, `Si votre voyage se passe comme prévu, vous devriez bientôt arriver à destination.`},
	{`आगमन समय`, `यदि आपके यात्रा का प्रबंधन समयानुसार है तो आप अपनी मंजिल पर पहुंच जाएंगे।`},
	{`到着時間`, `旅程通りに進行していれば、間もなく目的地に到着するはずです。`},
	{`Horário de Chegada`, `Se os seus planos de viagem estão em execução no momento, você deve estar chegando ao seu destino.`},
	{`Время прибытия`, `Если все ваши рейсы идут по расписанию, вы должны быть рядом с пунктом назначения.`},
	{`抵达时间`, `如果出行安排在准时推进，那么您应当很快就会抵达目的地。`},
}
