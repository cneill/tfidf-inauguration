# Inaugural speech analyzer

## About

This computes the [tf-idf](http://www.tfidf.com/) on presidential inauguration
speeches. The speech texts are included with the repo.

This dataset only covers presidents Ronald Reagan, George H.W. Bush,
Bill Clinton, George W. Bush, Barack Obama, and Donald Trump

The speeches were retrived [here](http://avalon.law.yale.edu/subject_menus/inaug.asp)
and [here](http://www.presidency.ucsb.edu/inaugurals.php).

## Using the tool

```
go build
./inauguration
```

## Data

```
Speech: speeches/barack-obama-1.txt; Words: 2391
TFIDF:
	carried: 0.002889
	false: 0.001926
	faced: 0.001926
	calls: 0.001926
	lower: 0.001926
	icy: 0.001926
	waters: 0.001926
	virtue: 0.001926
	charter: 0.001926
	father: 0.001378
	conflict: 0.001346

Speech: speeches/barack-obama-2.txt; Words: 2090
TFIDF:
	complete: 0.003850
	requires: 0.003850
	she: 0.003080
	until: 0.002628
	knowing: 0.002310
	lessons: 0.002203
	generation's: 0.002203
	train: 0.002203
	king: 0.002203
	initiative: 0.002203
	harm: 0.002203

Speech: speeches/bill-clinton-1.txt; Words: 1583
TFIDF:
	season: 0.005818
	renewal: 0.004067
	whom: 0.004067
	change: 0.003503
	posterity: 0.003050
	spring: 0.003050
	raised: 0.003050
	sake: 0.003050
	serving: 0.002909
	compete: 0.002909
	idea: 0.002189

Speech: speeches/bill-clinton-2.txt; Words: 2157
TFIDF:
	enough: 0.003056
	20th: 0.002238
	21st: 0.002238
	awful: 0.002135
	moved: 0.002135
	waste: 0.002135
	labors: 0.002135
	quest: 0.002135
	seemed: 0.002135
	perfect: 0.002135
	19th: 0.002135

Speech: speeches/donald-trump-1.txt; Words: 1439
TFIDF:
	protected: 0.005592
	obama: 0.004800
	countries: 0.004800
	we've: 0.004800
	exists: 0.003200
	industry: 0.003200
	breath: 0.003200
	transferring: 0.003200
	winning: 0.003200
	glorious: 0.003200
	politicians: 0.003200

Speech: speeches/george-hw-bush.txt; Words: 2317
TFIDF:
	breeze: 0.004969
	word: 0.004969
	door: 0.003975
	fact: 0.003473
	loyal: 0.002981
	blowing: 0.002981
	expression: 0.002981
	other's: 0.002981
	don't: 0.002845
	friends: 0.002393
	engagement: 0.002084

Speech: speeches/george-w-bush-1.txt; Words: 1579
TFIDF:
	story: 0.006262
	civility: 0.005833
	affirm: 0.003058
	laws: 0.002917
	angel: 0.002917
	stakes: 0.002917
	rides: 0.002917
	grand: 0.002917
	whirlwind: 0.002917
	directs: 0.002917
	sometimes: 0.002783

Speech: speeches/george-w-bush-2.txt; Words: 2054
TFIDF:
	excuse: 0.003363
	fire: 0.003134
	defended: 0.002351
	questions: 0.002351
	goal: 0.002351
	institutions: 0.002351
	permanent: 0.002351
	ideal: 0.002351
	proclaimed: 0.002242
	resentment: 0.002242
	fulfill: 0.002242

Speech: speeches/ronald-reagan-1.txt; Words: 2438
TFIDF:
	productivity: 0.002833
	having: 0.002833
	maintaining: 0.002833
	burden: 0.002833
	intention: 0.002833
	weapon: 0.002833
	special: 0.002641
	heroes: 0.002253
	present: 0.001980
	why: 0.001980
	front: 0.001980

Speech: speeches/ronald-reagan-2.txt; Words: 2563
TFIDF:
	increase: 0.003594
	government's: 0.002695
	song: 0.002695
	weapons: 0.002572
	nuclear: 0.002572
	human: 0.002434
	reduce: 0.002143
	federal: 0.002143
	sound: 0.001884
	spend: 0.001884
	destroy: 0.001884
```
