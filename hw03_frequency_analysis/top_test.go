package hw03_frequency_analysis //nolint:golint
import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Change to true if needed
var taskWithAsteriskIsCompleted = true

var text = `Как ,видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`
var specail1 = "~`!@ #$% ^&*() +{  }[]|;':\",.<>?"
var digits = ` 1234 324 234
	234234 324234234234234 1234 1234 1 2 3 1 1
	1 1 1 1 1 1`
var postfix = "a, a! a\\ a b^"
var one = "  ddd  "
var dash = "aaabbb, aaa-bbb"

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		assert.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			fmt.Println(Top10(text))
			expected := []string{"он", "а", "и", "что", "ты", "не", "если", "то", "его", "кристофер", "робин", "в"}
			assert.Subset(t, expected, Top10(text))
		} else {
			expected := []string{"он", "и", "а", "что", "ты", "не", "если", "-", "то", "Кристофер"}
			assert.ElementsMatch(t, expected, Top10(text))
		}
	})
	t.Run("special symbols", func(t *testing.T) {
		assert.Len(t, Top10(specail1), 0)
	})
	t.Run("digits", func(t *testing.T) {
		expected := []string{"1", "1234", "2", "3", "234234", "324234234234234", "324", "234"}
		assert.ElementsMatch(t, expected, Top10(digits))
	})
	t.Run("different postfix", func(t *testing.T) {
		expected := []string{"a", "b"}
		assert.ElementsMatch(t, expected, Top10(postfix))
	})
	t.Run("one word", func(t *testing.T) {
		expected := []string{"ddd"}
		assert.ElementsMatch(t, expected, Top10(one))
	})
	t.Run("dash inside", func(t *testing.T) {
		expected := []string{"aaabbb", "aaa-bbb"}
		assert.ElementsMatch(t, expected, Top10(dash))
	})
}
