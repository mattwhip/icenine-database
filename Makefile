TEST_REPORTS="admin" "bot" "daily_bonus" "matchmaking_processor" "user_accounts" "user_data"

test:
	mkdir -p ./test-reports
	ginkgo -r

	for testDir in ${TEST_REPORTS} ; do \
		mv ./$$testDir/test*.xml ./test-reports ; \
	done
