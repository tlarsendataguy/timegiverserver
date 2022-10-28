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

var StopCaffeine = []StepText{
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

var Caffeine12 = []StepText{
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
